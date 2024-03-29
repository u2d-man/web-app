package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db = NewDBConn()

func NewDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	db.LogMode(true)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetDBConn() *gorm.DB {
	return db
}

func GetDBConfig() (string, string) {
	return "mysql", "root:root@tcp(127.0.0.1)/test_database?charset=utf8&parseTime=True&loc=Asia%2FTokyo"
}
