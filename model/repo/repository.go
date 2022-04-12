package repo

import (
	"gomockt/domain"
	"gomockt/model/db/tx"
)

type BookRepository interface {
	tx.TxDataInterface

	QueryBook(title string) (result []domain.Book, err error)
	InsertBook(book domain.Book) (err error)
	UpdateBook(book domain.Book) (err error)
}
