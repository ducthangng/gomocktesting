package mysql

import (
	"gomockt/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertBook(t *testing.T) {
	q := NewQuerierTesting()

	err := q.InsertBook(domain.Book{
		Title: "hello world 2",
	})

	require.NoError(t, err)
}

func TestQueryBook(t *testing.T) {
	q := NewQuerierTesting()

	book, err := q.QueryBook("hello world")

	require.NoError(t, err)
	require.Equal(t, "hello world", book[0].Title)
}
