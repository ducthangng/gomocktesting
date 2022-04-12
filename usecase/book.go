package usecase

import (
	"gomockt/domain"
	"gomockt/model/repo"
)

type BookInteractor struct {
	repo repo.BookRepository
}

func NewBookInteractor(repo repo.BookRepository) *BookInteractor {
	return &BookInteractor{
		repo: repo,
	}
}

// If there already exist a book with the same title
// Then change that existed book into "${book.Title} 1" and current insert book into "${book.Title} 2"
func (b *BookInteractor) InsertBook(book domain.Book) (err error) {
	existedBook, err := b.repo.QueryBook(book.Title)
	if err != nil {
		return err
	}

	if existedBook != nil {
		existedBook[0].Title = existedBook[0].Title + " 1"
		err = b.repo.UpdateBook(existedBook[0])
		if err != nil {
			b.repo.EnableTx(err)
			return err
		}

		book.Title = book.Title + " 2"
	}

	if err := b.repo.InsertBook(book); err != nil {
		b.repo.EnableTx(err)
		return err
	}

	return b.repo.EnableTx(nil)
}
