package postgres

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/ssentinull/create-apis-using-golang/internal/model"
	"github.com/ssentinull/create-apis-using-golang/internal/utils"
	"gorm.io/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) model.BookRepository {
	return &bookRepo{db: db}
}

func (br *bookRepo) Create(ctx context.Context, book *model.Book) error {
	err := br.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(book).Error; err != nil {
			return err

		}
		return nil
	})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":  utils.Dump(ctx),
			"book": utils.Dump(book),
		}).Error(err)
		return err
	}

	return nil
}

func (br *bookRepo) DeleteByID(ctx context.Context, ID int64) error {
	err := br.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.Book{}, ID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx": utils.Dump(ctx),
			"ID":  ID,
		}).Error(err)
		return err
	}

	return nil
}

func (br *bookRepo) FindByID(ctx context.Context, ID int64) (*model.Book, error) {
	book := &model.Book{}
	err := br.db.WithContext(ctx).Where("id = ?", ID).Take(&book).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx": utils.Dump(ctx),
			"ID":  ID,
		}).Error(err)
		return nil, err
	}

	return book, nil
}

func (br *bookRepo) FindAll(ctx context.Context) ([]*model.Book, error) {
	books := []*model.Book{}
	err := br.db.WithContext(ctx).Order("id DESC").Find(&books).Error
	if err != nil {
		logrus.WithField("ctx", utils.Dump(ctx)).Error(err)
		return nil, err
	}

	return books, nil
}

func (br *bookRepo) Update(ctx context.Context, book *model.Book) (*model.Book, error) {
	err := br.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Updates(book).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":  utils.Dump(ctx),
			"book": utils.Dump(book),
		}).Error(err)
		return nil, err
	}

	return br.FindByID(ctx, book.ID)
}
