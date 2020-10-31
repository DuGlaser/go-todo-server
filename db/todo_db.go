package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "username"
	password = "password"
	host     = "127.0.0.1:3306"
	schema   = "db_name"
)

var TodoDB todoDBInterface = &todoDB{}

type todoDBInterface interface{}

type todoDB struct {
	client *sql.DB
}

func (db *todoDB) init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	var err error
	db.client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = db.client.Ping(); err != nil {
		panic(err)
	}
}
