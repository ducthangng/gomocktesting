package mysql

import (
	"context"
	"gomockt/domain"
	"log"
)

func (q *Querier) QueryBook(title string) (result []domain.Book, err error) {
	rows, err := q.DB.QueryContext(context.Background(), "SELECT * FROM books WHERE title = ?", title)
	if err != nil {
		log.Println("failed with error: ", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var book domain.Book
		err = rows.Scan(&book.ID, &book.Title)
		if err != nil {
			log.Println("failed with error: ", err)
			return
		}
		result = append(result, book)
	}

	return result, nil
}

func (q *Querier) InsertBook(book domain.Book) (err error) {
	// return errors.New("not implemented")
	_, err = q.DB.ExecContext(context.Background(), "INSERT INTO books SET title = ?", book.Title)
	if err != nil {
		log.Println("failed with error: ", err)
		return
	}

	return nil
}

func (q *Querier) UpdateBook(book domain.Book) (err error) {
	// return errors.New("not implemented")
	_, err = q.DB.ExecContext(context.Background(), "UPDATE books SET title = ? WHERE id = ?", book.Title, book.ID)
	if err != nil {
		log.Println("failed with error: ", err)
		return
	}

	return nil
}
