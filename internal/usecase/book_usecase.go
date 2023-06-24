package usecase

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/ssentinull/create-apis-using-golang/internal/model"
	"github.com/ssentinull/create-apis-using-golang/internal/utils"
)

type bookUsecase struct {
	bookRepo model.BookRepository
}

func NewBookUsecase(br model.BookRepository) model.BookUsecase {
	return &bookUsecase{bookRepo: br}
}

func (bu *bookUsecase) Create(ctx context.Context, book *model.Book) (*model.Book, error) {
	if err := bu.bookRepo.Create(ctx, book); err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":  utils.Dump(ctx),
			"book": utils.Dump(book),
		}).Error(err)
		return nil, err
	}

	return book, nil
}

func (bu *bookUsecase) DeleteByID(ctx context.Context, ID int64) error {
	if err := bu.bookRepo.DeleteByID(ctx, ID); err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx": utils.Dump(ctx),
			"ID":  ID,
		}).Error(err)
		return err
	}

	return nil
}

func (bu *bookUsecase) FindByID(ctx context.Context, ID int64) (*model.Book, error) {
	book, err := bu.bookRepo.FindByID(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx": utils.Dump(ctx),
			"ID":  ID,
		}).Error(err)
		return nil, err
	}

	return book, nil
}

func (bu *bookUsecase) FindAll(ctx context.Context, params model.GetBooksQueryParams) ([]*model.Book, int64, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":    utils.Dump(ctx),
		"params": utils.Dump(params),
	})

	books, err := bu.bookRepo.FindAll(ctx, params)
	if err != nil {
		logger.Error(err)
		return nil, int64(0), err
	}

	count, err := bu.bookRepo.CountAll(ctx)
	if err != nil {
		logger.Error(err)
		return nil, int64(0), err
	}

	return books, count, nil
}

func (bu *bookUsecase) Update(ctx context.Context, book *model.Book) (*model.Book, error) {
	book, err := bu.bookRepo.Update(ctx, book)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":  utils.Dump(ctx),
			"book": utils.Dump(book),
		}).Error(err)
		return nil, err
	}

	return book, nil
}
