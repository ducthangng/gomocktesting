package mysql

import (
	"database/sql"
	"gomockt/config"
	"gomockt/model/db/sqlcon"
	"gomockt/model/db/tx"
	"log"
)

func NewQuerierTesting() *Querier {
	config, err := config.ReadConfig("../../../")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlcon.OpenConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	conn := BuildTx(db, false)

	return NewQuerier(conn)
}

func BuildTx(db *sql.DB, Flag bool) tx.DBTX {
	if Flag {
		txs, err := db.Begin()
		if err != nil {
			return nil
		}
		return &tx.TxConn{DB: txs}
	}

	return &tx.DbConn{DB: db}
}
