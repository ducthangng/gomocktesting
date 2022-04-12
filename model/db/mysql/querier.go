package mysql

import (
	"gomockt/model/db/tx"
)

// Querier implements the repo interface
type Querier struct {
	DB tx.DBTX
}

func NewQuerier(DB tx.DBTX) *Querier {
	return &Querier{
		DB: DB,
	}
}

func (q *Querier) EnableTx(txErr error) error {
	return q.DB.TxEnd(txErr)
}
