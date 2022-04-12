package mysql

import "gomockt/model/db/tx"

// Querier implements the repo interface
type Querier struct {
	DB tx.DBTX
}

// EnableTx allow the transaction to be initialized.
func (q *Querier) EnableTx(txFunc func() error) error {
	return q.DB.TxEnd(txFunc)
}

func NewQuerier(DB tx.DBTX) *Querier {
	return &Querier{
		DB: DB,
	}
}
