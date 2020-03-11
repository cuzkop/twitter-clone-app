package models

import (
	"database/sql"
	"log"
)

type DB struct {
	DB *sql.DB
}

func init() {
	log.SetPrefix("[db]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func NewSqlHandler() *DB {
	conn, err := sql.Open("mysql", "root@tcp(127.0.0.1)/twitter")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(DB)
	sqlHandler.DB = conn
	return sqlHandler
}
