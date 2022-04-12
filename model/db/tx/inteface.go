package tx

// TxDataInterface represents operations needed for transaction support.
// It only needs to be implemented once for each database
type TxDataInterface interface {
	// EnableTx is called at the end of a transaction and based on whether there is an error, it commits or rollback the
	// transaction.
	// txFunc is the business function wrapped in a transaction
	EnableTx(txErr error) error
}
