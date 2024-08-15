package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"net/http"
)

type Text struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
}

func InitMysql() *gorm.DB {
	var err error
	engine, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s)/powervoting?charset=utf8&parseTime=True&loc=Local", "root", "Csystem32...", "localhost"),
		DefaultStringSize:         256,   // string default length
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbl_", // 表前缀增加tbl_
			SingularTable: true,   // 单数形式，User结构体对应user表，非users
		},
	})
	if err != nil {
		log.Fatalln("连接数据库异常：", err)
	}
	return engine
}

func main() {

	mysqlEngine := InitMysql()

}
