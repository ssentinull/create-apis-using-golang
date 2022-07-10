package model

import (
	"context"
	"time"
)

type Book struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Description   string    `json:"description"`
	PublishedDate string    `json:"published_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

type CreateBookInput struct {
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Description   string    `json:"description"`
	PublishedDate string    `json:"published_date"`
	CreatedAt     time.Time `json:"created_at"`
}

func (i CreateBookInput) ToModel() *Book {
	return &Book{
		ID:            int64(1),
		Title:         i.Title,
		Author:        i.Author,
		Description:   i.Description,
		PublishedDate: i.PublishedDate,
		CreatedAt:     i.CreatedAt,
	}
}

type BookUsecase interface {
	Create(ctx context.Context, input *CreateBookInput) (book *Book, err error)
	FindByID(ctx context.Context, ID int64) (book *Book, err error)
	FindAll(ctx context.Context) (books []*Book, err error)
}

type BookRepository interface {
	Create(ctx context.Context, book *Book) (err error)
	FindByID(ctx context.Context, ID int64) (book *Book, err error)
	FindAll(ctx context.Context) (books []*Book, err error)
}
