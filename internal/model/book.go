package model

import (
	"context"
	"time"

	"github.com/ssentinull/create-apis-using-golang/internal/utils"
	"gorm.io/gorm"
)

type Book struct {
	ID            int64          `json:"id"`
	Title         string         `json:"title"`
	Author        string         `json:"author"`
	Description   string         `json:"description"`
	PublishedDate string         `json:"published_date"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

type CreateBookInput struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Description   string `json:"description"`
	PublishedDate string `json:"published_date"`
}

func (i CreateBookInput) ToModel() *Book {
	return &Book{
		ID:            utils.GenerateID(),
		Title:         i.Title,
		Author:        i.Author,
		Description:   i.Description,
		PublishedDate: i.PublishedDate,
		CreatedAt:     time.Now(),
	}
}

type UpdateBookInput struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Description   string `json:"description"`
	PublishedDate string `json:"published_date"`
}

func (i UpdateBookInput) ToModel() *Book {
	return &Book{
		ID:            i.ID,
		Title:         i.Title,
		Author:        i.Author,
		Description:   i.Description,
		PublishedDate: i.PublishedDate,
		UpdatedAt:     time.Now(),
	}
}

type BookUsecase interface {
	Create(ctx context.Context, input *Book) (book *Book, err error)
	DeleteByID(ctx context.Context, ID int64) (err error)
	FindByID(ctx context.Context, ID int64) (book *Book, err error)
	FindAll(ctx context.Context) (books []*Book, err error)
	Update(ctx context.Context, input *Book) (book *Book, err error)
}

type BookRepository interface {
	Create(ctx context.Context, input *Book) (err error)
	DeleteByID(ctx context.Context, ID int64) (err error)
	FindByID(ctx context.Context, ID int64) (book *Book, err error)
	FindAll(ctx context.Context) (books []*Book, err error)
	Update(ctx context.Context, input *Book) (book *Book, err error)
}
