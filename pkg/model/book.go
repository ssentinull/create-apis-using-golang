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
	UpdatedAt     time.Time `json:"updated_at"`
}

func (i CreateBookInput) ToModel() *Book {
	return &Book{
		ID:            int64(1),
		Title:         i.Title,
		Author:        i.Author,
		Description:   i.Description,
		PublishedDate: i.PublishedDate,
		CreatedAt:     i.CreatedAt,
		UpdatedAt:     i.UpdatedAt,
	}
}

type BookUsecase interface {
	CreateBook(context.Context, *CreateBookInput) (*Book, error)
	GetBookByID(context.Context, int64) (Book, error)
	GetBooks(context.Context) ([]Book, error)
}

type BookRepository interface {
	CreateBook(context.Context, *Book) error
	ReadBookByID(context.Context, int64) (Book, error)
	ReadBooks(context.Context) ([]Book, error)
}
