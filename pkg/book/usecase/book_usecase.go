package usecase

import (
	"context"
	"github.com/ssentinull/create-apis-using-golang/pkg/utils"

	"github.com/sirupsen/logrus"
	"github.com/ssentinull/create-apis-using-golang/pkg/model"
)

type bookUsecase struct {
	bookRepo model.BookRepository
}

func NewBookUsecase(br model.BookRepository) model.BookUsecase {
	return &bookUsecase{bookRepo: br}
}

func (bu *bookUsecase) Create(ctx context.Context, input *model.CreateBookInput) (*model.Book, error) {
	book := input.ToModel()
	if err := bu.bookRepo.Create(ctx, book); err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":   utils.Dump(ctx),
			"input": utils.Dump(input),
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

func (bu *bookUsecase) FindAll(ctx context.Context) ([]*model.Book, error) {
	books, err := bu.bookRepo.FindAll(ctx)
	if err != nil {
		logrus.WithField("ctx", utils.Dump(ctx)).Error(err)
		return nil, err
	}

	return books, nil
}

func (bu *bookUsecase) Update(ctx context.Context, input *model.UpdateBookInput) (*model.Book, error) {
	book, err := bu.bookRepo.Update(ctx, input.ToModel())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":   utils.Dump(ctx),
			"input": utils.Dump(input),
		}).Error(err)
		return nil, err
	}

	return book, nil
}
