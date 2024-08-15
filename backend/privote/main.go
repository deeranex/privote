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
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbl_",
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalln("Database connect error：", err)
	}
	return engine
}

func main() {

	mysqlEngine := InitMysql()

}
