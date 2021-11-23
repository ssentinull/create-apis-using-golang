package model

import (
	"context"
	"time"
)

type Book struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type BookUsecase interface {
	GetBookByID(context.Context, int64) (Book, error)
	GetBooks(context.Context) ([]Book, error)
}

type BookRepository interface {
	ReadBookByID(context.Context, int64) (Book, error)
	ReadBooks(context.Context) ([]Book, error)
}
