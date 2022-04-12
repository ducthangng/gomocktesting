package sqlcon

import (
	"database/sql"
	"fmt"
	"gomockt/config"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection(config config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.DatabaseHost,
		config.Name,
	))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
