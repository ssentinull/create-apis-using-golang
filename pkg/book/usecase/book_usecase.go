package usecase

import (
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/ssentinull/create-apis-using-golang/pkg/model"
)

type bookUsecase struct {
	bookRepo model.BookRepository
}

func NewBookUsecase(br model.BookRepository) model.BookUsecase {
	return &bookUsecase{bookRepo: br}
}

func (bu *bookUsecase) CreateBook(ctx context.Context, input *model.CreateBookInput) (*model.Book, error) {
	book := input.ToModel()
	if err := bu.bookRepo.CreateBook(ctx, book); err != nil {
		return nil, err
	}

	return book, nil
}

func (bu *bookUsecase) GetBookByID(ctx context.Context, ID int64) (model.Book, error) {
	book, err := bu.bookRepo.ReadBookByID(ctx, ID)
	if err != nil {
		c, err := json.Marshal(ctx)
		if err != nil {
			logrus.Error(err)
		}

		logrus.WithFields(logrus.Fields{
			"ctx": c,
			"ID":  ID,
		}).Error(err)

		return model.Book{}, err
	}

	return book, nil
}

func (bu *bookUsecase) GetBooks(ctx context.Context) ([]model.Book, error) {
	books, err := bu.bookRepo.ReadBooks(ctx)
	if err != nil {
		c, err := json.Marshal(ctx)
		if err != nil {
			logrus.Error(err)
		}

		logrus.WithField("ctx", c).Error(err)

		return nil, err
	}

	return books, nil
}
