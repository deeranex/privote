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

	engine := gin.Default()
	engine.POST("update", func(c *gin.Context) {
        var text Text
        err := c.BindJSON(&text)
        if err != nil {
            log.Println("参数异常")
            c.JSON(http.StatusBadRequest, gin.H{
                "msg": "参数异常",
            })
        }
        tx := mysqlEngine.Table("tbl_text").Create(&text)
        if tx.Error != nil {
            log.Println("create error", tx.Error.Error())
        }
        c.JSON(http.StatusOK, gin.H{
            "msg": "OK",
            "id":  text.ID,
        })
    })

	engine.GET("/get/:id", func(c *gin.Context) {
        id := c.Param("id")
        var text Text
        tx := mysqlEngine.Table("tbl_text").Where("id = ?", id).Find(&text)
        if tx.Error != nil {
            log.Println("get error:", tx.Error.Error())
            c.JSON(http.StatusBadRequest, gin.H{
                "msg": "get text error",
            })
        }
        c.JSON(http.StatusOK, gin.H{
            "msg":  "OK",
            "data": text,
        })
    })

    engine.Run(":9999")
}
