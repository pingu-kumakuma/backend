package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func DbConnect() {
	dsn := "test_user:password@tcp(localhost:3308)/test_database"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
}
