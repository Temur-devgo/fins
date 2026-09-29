package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(driver string, connection string) {
	var err error
	db, err = sql.Open(driver, connection)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxOpenConns(250)
	db.SetConnMaxLifetime(5 * time.Minute)
}
