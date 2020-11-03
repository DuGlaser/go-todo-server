package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
}

func NewDB() *sql.DB {
	c := NewTodoConfig()

	return newDB(&DB{
		Host:     c.DB.Production.Host,
		Username: c.DB.Production.Username,
		Password: c.DB.Production.Password,
		DBName:   c.DB.Production.DBName,
		Port:     c.DB.Production.Port,
	})
}

func NewTestDB() *sql.DB {
	c := NewTodoConfig()

	return newDB(&DB{
		Host:     c.DB.Test.Host,
		Username: c.DB.Test.Username,
		Password: c.DB.Test.Password,
		DBName:   c.DB.Test.DBName,
		Port:     c.DB.Test.Port,
	})
}

func newDB(d *DB) *sql.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.DBName,
	)

	DBConn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = DBConn.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connect mysql!")
	return DBConn
}
