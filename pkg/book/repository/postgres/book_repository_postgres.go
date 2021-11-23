package postgres

import (
	"context"
	"time"

	"github.com/ssentinull/create-apis-using-golang/pkg/model"
)

type bookRepo struct{}

func NewBookRepository() model.BookRepository {
	return &bookRepo{}
}

func (br *bookRepo) ReadBookByID(ctx context.Context, ID int64) (model.Book, error) {
	book := model.Book{
		ID:          ID,
		Title:       "Harry Potter",
		Author:      "J. K. Rowling",
		Description: "A book about wizards",
		PublishedAt: time.Now(),
		CreatedAt:   time.Now(),
	}

	return book, nil
}

func (br *bookRepo) ReadBooks(ctx context.Context) ([]model.Book, error) {
	books := []model.Book{
		{
			ID:          1,
			Title:       "Harry Potter",
			Author:      "J. K. Rowling",
			Description: "A book about wizards",
			PublishedAt: time.Now(),
			CreatedAt:   time.Now(),
		},
		{
			ID:          2,
			Title:       "The Hobbit",
			Author:      "J. R. R. Tolkien",
			Description: "A book about hobbits",
			PublishedAt: time.Now(),
			CreatedAt:   time.Now(),
		},
	}

	return books, nil
}
