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

var dbConn *sql.DB

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	var err error
	dbConn, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = dbConn.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connect mysql!")
}
