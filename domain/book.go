package domain

type Book struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}
