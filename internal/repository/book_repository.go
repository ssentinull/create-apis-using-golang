package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/ssentinull/create-apis-using-golang/internal/model"
	"github.com/ssentinull/create-apis-using-golang/internal/utils"
	"gorm.io/gorm"
)

type bookRepo struct {
	db        *gorm.DB
	cacheRepo model.CacheRepository
}

func NewBookRepository(db *gorm.DB, cacheRepo model.CacheRepository) model.BookRepository {
	return &bookRepo{
		db:        db,
		cacheRepo: cacheRepo,
	}
}

func (br *bookRepo) Create(ctx context.Context, book *model.Book) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":  utils.Dump(ctx),
		"book": utils.Dump(book),
	})

	err := br.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(book).Error; err != nil {
			return err

		}
		return nil
	})

	if err != nil {
		logger.Error(err)
		return err
	}

	cacheKeys := []string{
		br.cacheHash(),
		br.countAllCacheKey(),
	}

	if err := br.cacheRepo.Delete(ctx, cacheKeys...); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (br *bookRepo) DeleteByID(ctx context.Context, ID int64) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.Dump(ctx),
		"ID":  ID,
	})

	err := br.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.Book{}, ID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		logger.Error(err)
		return err
	}

	cacheKeys := []string{
		br.findByIDCacheKey(ID),
		br.cacheHash(),
	}

	if err := br.cacheRepo.Delete(ctx, cacheKeys...); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (br *bookRepo) FindByID(ctx context.Context, ID int64) (*model.Book, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.Dump(ctx),
		"ID":  ID,
	})

	cacheKey := br.findByIDCacheKey(ID)
	reply, err := br.cacheRepo.Get(ctx, cacheKey)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if reply != "" {
		book := &model.Book{}
		if err := json.Unmarshal([]byte(reply), &book); err != nil {
			logger.Error(err)
			return nil, err
		}
		return book, nil
	}

	book := &model.Book{}
	err = br.db.WithContext(ctx).Where("id = ?", ID).Take(&book).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	bytes, err := json.Marshal(book)
	if err != nil {
		logger.Error(err)
		return book, nil
	}

	if err := br.cacheRepo.Set(ctx, cacheKey, string(bytes)); err != nil {
		logger.Error(err)
	}

	return book, nil
}

func (br *bookRepo) FindAll(ctx context.Context, query model.GetBooksQueryParams) ([]*model.Book, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":   utils.Dump(ctx),
		"query": utils.Dump(query),
	})

	cacheHash := br.cacheHash()
	cacheKey := br.findAllByQueryParams(query)
	reply, err := br.cacheRepo.HashGet(ctx, cacheHash, cacheKey)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if reply != "" {
		books := []*model.Book{}
		if err := json.Unmarshal([]byte(reply), &books); err != nil {
			logger.Error(err)
			return nil, err
		}
		return books, nil
	}

	books := []*model.Book{}
	err = br.db.WithContext(ctx).
		Order("id DESC").
		Offset(int(model.Offset(query.Page, query.Size))).
		Limit(int(query.Size)).
		Find(&books).
		Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	bytes, err := json.Marshal(books)
	if err != nil {
		logger.Error(err)
		return books, nil
	}

	if err := br.cacheRepo.HashSet(ctx, cacheHash, cacheKey, string(bytes)); err != nil {
		logger.Error(err)
	}

	return books, nil
}

func (br *bookRepo) CountAll(ctx context.Context) (int64, error) {
	logger := logrus.WithField("ctx", utils.Dump(ctx))

	cacheKey := br.countAllCacheKey()
	reply, err := br.cacheRepo.Get(ctx, cacheKey)
	if err != nil {
		logger.Error(err)
		return 0, err
	}

	if reply != "" {
		count := int64(0)
		if err := json.Unmarshal([]byte(reply), &count); err != nil {
			logger.Error(err)
			return 0, err
		}
		return count, nil
	}

	count := int64(0)
	err = br.db.WithContext(ctx).
		Model(model.Book{}).
		Count(&count).
		Error
	if err != nil {
		logrus.WithField("ctx", utils.Dump(ctx)).Error(err)
		return int64(0), err
	}

	bytes, err := json.Marshal(count)
	if err != nil {
		logger.Error(err)
		return 0, err
	}

	if err := br.cacheRepo.Set(ctx, cacheKey, string(bytes)); err != nil {
		logger.Error(err)
	}

	return count, nil
}

func (br *bookRepo) Update(ctx context.Context, book *model.Book) (*model.Book, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":  utils.Dump(ctx),
		"book": utils.Dump(book),
	})

	err := br.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Updates(book).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	cacheKeys := []string{
		br.cacheHash(),
		br.countAllCacheKey(),
		br.findByIDCacheKey(book.ID),
	}

	if err := br.cacheRepo.Delete(ctx, cacheKeys...); err != nil {
		logger.Error(err)
		return nil, err
	}

	return br.FindByID(ctx, book.ID)
}

func (br *bookRepo) cacheHash() string {
	return "book"
}

func (br *bookRepo) findByIDCacheKey(ID int64) string {
	return fmt.Sprintf("book:%d", ID)
}

func (br *bookRepo) findAllByQueryParams(query model.GetBooksQueryParams) string {
	return fmt.Sprintf("book:page:%d:size:%d", query.Page, query.Size)
}

func (br *bookRepo) countAllCacheKey() string {
	return "book:count"
}
