package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:ducthang@tcp(127.0.0.1:3306)/books?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
