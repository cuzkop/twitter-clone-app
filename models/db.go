package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	DB *gorm.DB
}

var sqlHandler *DB

func init() {
	log.SetPrefix("[models/db]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func NewSqlHandler() {
	conn, err := gorm.Open("mysql", "root:PASSWORD@tcp(127.0.0.1)/twitter?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	sqlHandler = new(DB)
	sqlHandler.DB = conn
}

func GetSqlHandler() *DB {
	if sqlHandler == nil {
		NewSqlHandler()
	}

	return sqlHandler
}
