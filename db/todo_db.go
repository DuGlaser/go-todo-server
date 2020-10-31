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

type todoDBInterface interface {
	GetClient() *sql.DB
	Init()
}

type todoDB struct {
	Client *sql.DB
}

func (db *todoDB) GetClient() *sql.DB {
	return db.Client
}

func (db *todoDB) Init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	var err error
	db.Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = db.Client.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connect mysql")
}
