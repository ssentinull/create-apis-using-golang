package repository

import (
	"encoding/json"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/ssentinull/create-apis-using-golang/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestBookRepository_Create(t *testing.T) {
	mockedDependency := newMockedDependency(t)
	defer mockedDependency.close()

	ctx := mockedDependency.ctx
	repo := bookRepo{
		db:        mockedDependency.db,
		cacheRepo: mockedDependency.cacheRepo,
	}

	book := model.Book{
		ID:          int64(1),
		Title:       "Harry Potter",
		Author:      "J. K. Rowling",
		Description: "A series about wizards",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	cacheKeys := []string{
		repo.cacheHash(),
		repo.countAllCacheKey(),
	}

	query := `INSERT INTO "books" ("title","author","description","published_date","created_at","updated_at","deleted_at","id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING "id"`

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "author", "title", "description"}).
			AddRow(book.ID, book.Author, book.Title, book.Description)

		mockedDependency.sql.ExpectBegin()
		mockedDependency.sql.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
		mockedDependency.sql.ExpectCommit()
		mockedDependency.cacheRepo.EXPECT().Delete(ctx, cacheKeys).Times(1).Return(nil)

		err := repo.Create(ctx, &book)
		assert.NoError(t, err)
	})

	t.Run("failed - create book in db return error", func(t *testing.T) {
		mockedDependency.sql.ExpectBegin()
		mockedDependency.sql.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("db error"))
		mockedDependency.sql.ExpectRollback()

		err := repo.Create(ctx, &book)
		assert.Error(t, err)
	})

	t.Run("failed - delete cache return error", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "author", "title", "description"}).
			AddRow(book.ID, book.Author, book.Title, book.Description)

		mockedDependency.sql.ExpectBegin()
		mockedDependency.sql.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
		mockedDependency.sql.ExpectCommit()
		mockedDependency.cacheRepo.EXPECT().Delete(ctx, cacheKeys).Times(1).Return(errors.New("cache error"))

		err := repo.Create(ctx, &book)
		assert.Error(t, err)
	})
}

func TestBookRepository_Delete(t *testing.T) {
	mockedDependency := newMockedDependency(t)
	defer mockedDependency.close()

	ctx := mockedDependency.ctx
	repo := bookRepo{
		db:        mockedDependency.db,
		cacheRepo: mockedDependency.cacheRepo,
	}

	book := model.Book{
		ID:          int64(1),
		Title:       "Harry Potter",
		Author:      "J. K. Rowling",
		Description: "A series about wizards",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	cacheKeys := []string{
		repo.findByIDCacheKey(book.ID),
		repo.cacheHash(),
	}

	query := `UPDATE "books" SET "deleted_at"=$1 WHERE "books"."id" = $2 AND "books"."deleted_at" IS NULL`

	t.Run("success", func(t *testing.T) {
		mockedDependency.sql.ExpectBegin()
		mockedDependency.sql.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(1, 1))
		mockedDependency.sql.ExpectCommit()
		mockedDependency.cacheRepo.EXPECT().Delete(ctx, cacheKeys).Times(1).Return(nil)

		err := repo.DeleteByID(ctx, book.ID)
		assert.NoError(t, err)
	})

	t.Run("failed - create book in db return error", func(t *testing.T) {
		mockedDependency.sql.ExpectBegin()
		mockedDependency.sql.ExpectExec(regexp.QuoteMeta(query)).WillReturnError(errors.New("db error"))
		mockedDependency.sql.ExpectRollback()

		err := repo.DeleteByID(ctx, book.ID)
		assert.Error(t, err)
	})

	t.Run("failed - delete cache return error", func(t *testing.T) {
		mockedDependency.sql.ExpectBegin()
		mockedDependency.sql.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(1, 1))
		mockedDependency.sql.ExpectCommit()
		mockedDependency.cacheRepo.EXPECT().Delete(ctx, cacheKeys).Times(1).Return(errors.New("cache error"))

		err := repo.DeleteByID(ctx, book.ID)
		assert.Error(t, err)
	})
}

func TestBookRepository_FindByID(t *testing.T) {
	mockedDependency := newMockedDependency(t)
	defer mockedDependency.close()

	ctx := mockedDependency.ctx
	repo := bookRepo{
		db:        mockedDependency.db,
		cacheRepo: mockedDependency.cacheRepo,
	}

	book := model.Book{
		ID:          int64(1),
		Title:       "Harry Potter",
		Author:      "J. K. Rowling",
		Description: "A series about wizards",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	query := `SELECT * FROM "books" WHERE id = $1 AND "books"."deleted_at" IS NULL LIMIT 1`

	cacheKey := repo.findByIDCacheKey(book.ID)
	bytes, err := json.Marshal(book)
	assert.NoError(t, err)

	t.Run("success - fetch from cache", func(t *testing.T) {
		mockedDependency.cacheRepo.EXPECT().Get(ctx, cacheKey).Times(1).Return(string(bytes), nil)
		res, err := repo.FindByID(ctx, book.ID)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("success - fetch from db", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "author", "title", "description"}).
			AddRow(book.ID, book.Author, book.Title, book.Description)

		mockedDependency.cacheRepo.EXPECT().Get(ctx, cacheKey).Times(1).Return("", nil)
		mockedDependency.sql.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
		mockedDependency.cacheRepo.EXPECT().Set(ctx, cacheKey, gomock.Any()).Times(1).Return(nil)

		res, err := repo.FindByID(ctx, book.ID)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("failed - fetch from cache return error", func(t *testing.T) {
		mockedDependency.cacheRepo.EXPECT().Get(ctx, cacheKey).Times(1).Return("", errors.New("redis error"))
		res, err := repo.FindByID(ctx, book.ID)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("failed - fetch from db return error", func(t *testing.T) {
		mockedDependency.cacheRepo.EXPECT().Get(ctx, cacheKey).Times(1).Return("", nil)
		mockedDependency.sql.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("db error"))

		res, err := repo.FindByID(ctx, book.ID)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestBookRepository_FindAll(t *testing.T) {
	mockedDependency := newMockedDependency(t)
	defer mockedDependency.close()

	ctx := mockedDependency.ctx
	repo := bookRepo{
		db:        mockedDependency.db,
		cacheRepo: mockedDependency.cacheRepo,
	}

	queryParams := model.GetBooksQueryParams{
		Page: 1,
		Size: 5,
	}

	book := model.Book{
		ID:          int64(1),
		Title:       "Harry Potter",
		Author:      "J. K. Rowling",
		Description: "A series about wizards",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	query := `SELECT * FROM "books" WHERE "books"."deleted_at" IS NULL ORDER BY id DESC LIMIT 5`

	cacheHash := repo.cacheHash()
	cacheKey := repo.findAllByQueryParams(queryParams)

	books := []*model.Book{&book}
	bytes, err := json.Marshal(books)
	assert.NoError(t, err)

	t.Run("success - fetch from cache", func(t *testing.T) {
		mockedDependency.cacheRepo.EXPECT().HashGet(ctx, cacheHash, cacheKey).Times(1).Return(string(bytes), nil)
		res, err := repo.FindAll(ctx, queryParams)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("success - fetch from db", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "author", "title", "description"}).
			AddRow(book.ID, book.Author, book.Title, book.Description)

		mockedDependency.cacheRepo.EXPECT().HashGet(ctx, cacheHash, cacheKey).Times(1).Return("", nil)
		mockedDependency.sql.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
		mockedDependency.cacheRepo.EXPECT().HashSet(ctx, cacheHash, cacheKey, gomock.Any()).Times(1).Return(nil)

		res, err := repo.FindAll(ctx, queryParams)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("failed - fetch from cache return error", func(t *testing.T) {
		mockedDependency.cacheRepo.EXPECT().HashGet(ctx, cacheHash, cacheKey).Times(1).Return("", errors.New("redis error"))
		res, err := repo.FindAll(ctx, queryParams)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("failed - fetch from db return error", func(t *testing.T) {
		mockedDependency.cacheRepo.EXPECT().HashGet(ctx, cacheHash, cacheKey).Times(1).Return("", nil)
		mockedDependency.sql.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("db error"))

		res, err := repo.FindAll(ctx, queryParams)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestBookRepository_CountAll(t *testing.T) {
	mockedDependency := newMockedDependency(t)
	defer mockedDependency.close()

	ctx := mockedDependency.ctx
	repo := bookRepo{
		db:        mockedDependency.db,
		cacheRepo: mockedDependency.cacheRepo,
	}

	query := `SELECT count(*) FROM "books" WHERE "books"."deleted_at" IS NULL`

	cacheKey := repo.countAllCacheKey()
	bytes, err := json.Marshal(1)
	assert.NoError(t, err)

	t.Run("success - fetch from cache", func(t *testing.T) {
		mockedDependency.cacheRepo.EXPECT().Get(ctx, cacheKey).Times(1).Return(string(bytes), nil)
		res, err := repo.CountAll(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("success - fetch from db", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"count"}).AddRow(1)

		mockedDependency.cacheRepo.EXPECT().Get(ctx, cacheKey).Times(1).Return("", nil)
		mockedDependency.sql.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
		mockedDependency.cacheRepo.EXPECT().Set(ctx, cacheKey, gomock.Any()).Times(1).Return(nil)

		res, err := repo.CountAll(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("failed - fetch from cache return error", func(t *testing.T) {
		mockedDependency.cacheRepo.EXPECT().Get(ctx, cacheKey).Times(1).Return("", errors.New("redis error"))
		res, err := repo.CountAll(ctx)
		assert.Error(t, err)
		assert.Zero(t, res)
	})

	t.Run("failed - fetch from db return error", func(t *testing.T) {
		mockedDependency.cacheRepo.EXPECT().Get(ctx, cacheKey).Times(1).Return("", nil)
		mockedDependency.sql.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("db error"))

		res, err := repo.CountAll(ctx)
		assert.Error(t, err)
		assert.Zero(t, res)
	})
}

func TestBookRepository_Update(t *testing.T) {
	mockedDependency := newMockedDependency(t)
	defer mockedDependency.close()

	ctx := mockedDependency.ctx
	repo := bookRepo{
		db:        mockedDependency.db,
		cacheRepo: mockedDependency.cacheRepo,
	}

	book := model.Book{
		ID:          int64(1),
		Title:       "Harry Potter",
		Author:      "J. K. Rowling",
		Description: "A series about wizards",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	query := `UPDATE "books" SET "title"=$1,"author"=$2,"description"=$3,"updated_at"=$4 WHERE "books"."deleted_at" IS NULL AND "id" = $5`

	cacheKey := repo.findByIDCacheKey(book.ID)
	cacheKeys := []string{
		repo.cacheHash(),
		repo.countAllCacheKey(),
		repo.findByIDCacheKey(book.ID),
	}

	bytes, err := json.Marshal(book)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		mockedDependency.sql.ExpectBegin()
		mockedDependency.sql.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(1, 1))
		mockedDependency.sql.ExpectCommit()
		mockedDependency.cacheRepo.EXPECT().Delete(ctx, cacheKeys).Times(1).Return(nil)
		mockedDependency.cacheRepo.EXPECT().Get(ctx, cacheKey).Times(1).Return(string(bytes), nil)

		res, err := repo.Update(ctx, &book)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("failed - update book in db return error", func(t *testing.T) {
		mockedDependency.sql.ExpectBegin()
		mockedDependency.sql.ExpectExec(regexp.QuoteMeta(query)).WillReturnError(errors.New("db error"))
		mockedDependency.sql.ExpectRollback()

		res, err := repo.Update(ctx, &book)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("failed - delete cache return error", func(t *testing.T) {
		mockedDependency.sql.ExpectBegin()
		mockedDependency.sql.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(1, 1))
		mockedDependency.sql.ExpectCommit()
		mockedDependency.cacheRepo.EXPECT().Delete(ctx, cacheKeys).Times(1).Return(errors.New("cache error"))

		res, err := repo.Update(ctx, &book)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
