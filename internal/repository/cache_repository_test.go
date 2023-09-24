package repository

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ssentinull/create-apis-using-golang/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCacheRepository_Get(t *testing.T) {
	mockedDependency := newMockedDependency(t)
	defer mockedDependency.close()

	ctx := mockedDependency.ctx
	cacheRepo := cacheRepo{redisClient: mockedDependency.redis}
	book := model.Book{
		ID:          int64(1),
		Title:       "Harry Potter",
		Author:      "J. K. Rowling",
		Description: "A series about wizards",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	cacheKey := "book:1"
	cacheVal, err := json.Marshal(book)
	assert.NoError(t, err)

	t.Run("success - cache client return reply", func(t *testing.T) {
		mockedDependency.redisCmd.ExpectGet(cacheKey).SetVal(string(cacheVal))
		res, err := cacheRepo.Get(ctx, cacheKey)
		assert.NoError(t, err)
		assert.NotZero(t, res)
	})

	t.Run("success - cache client return empty string", func(t *testing.T) {
		mockedDependency.redisCmd.ExpectGet(cacheKey).SetErr(redis.Nil)
		res, err := cacheRepo.Get(ctx, cacheKey)
		assert.NoError(t, err)
		assert.Zero(t, res)
	})

	t.Run("failed - cache client return error", func(t *testing.T) {
		mockedDependency.redisCmd.ExpectGet(cacheKey).SetErr(errors.New("redis error"))
		res, err := cacheRepo.Get(ctx, cacheKey)
		assert.Error(t, err)
		assert.Zero(t, res)
	})
}

func TestCacheRepository_Set(t *testing.T) {
	mockedDependency := newMockedDependency(t)
	defer mockedDependency.close()

	ctx := mockedDependency.ctx
	cacheRepo := cacheRepo{redisClient: mockedDependency.redis}
	book := model.Book{
		ID:          int64(1),
		Title:       "Harry Potter",
		Author:      "J. K. Rowling",
		Description: "A series about wizards",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	cacheKey := "book:1"
	cacheVal, err := json.Marshal(book)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		mockedDependency.redisCmd.ExpectSet(cacheKey, string(cacheVal), 0).SetVal(string(cacheVal))
		err := cacheRepo.Set(ctx, cacheKey, string(cacheVal))
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		mockedDependency.redisCmd.ExpectSet(cacheKey, string(cacheVal), 0).SetErr(errors.New("redis error"))
		err := cacheRepo.Set(ctx, cacheKey, string(cacheVal))
		assert.Error(t, err)
	})
}

func TestCacheRepository_Delete(t *testing.T) {
	mockedDependency := newMockedDependency(t)
	defer mockedDependency.close()

	ctx := mockedDependency.ctx
	cacheRepo := cacheRepo{redisClient: mockedDependency.redis}
	cacheKey := "book:1"

	t.Run("success", func(t *testing.T) {
		mockedDependency.redisCmd.ExpectDel(cacheKey).SetVal(1)
		err := cacheRepo.Delete(ctx, cacheKey)
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		mockedDependency.redisCmd.ExpectDel(cacheKey).SetErr(errors.New("redis error"))
		err := cacheRepo.Delete(ctx, cacheKey)
		assert.Error(t, err)
	})
}

func TestCacheRepository_HashGet(t *testing.T) {
	mockedDependency := newMockedDependency(t)
	defer mockedDependency.close()

	ctx := mockedDependency.ctx
	cacheRepo := cacheRepo{redisClient: mockedDependency.redis}
	book := model.Book{
		ID:          int64(1),
		Title:       "Harry Potter",
		Author:      "J. K. Rowling",
		Description: "A series about wizards",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	cacheKey := "book:1"
	cacheHash := "hash:book"
	cacheVal, err := json.Marshal(book)
	assert.NoError(t, err)

	t.Run("success - cache client return reply", func(t *testing.T) {
		mockedDependency.redisCmd.ExpectHGet(cacheHash, cacheKey).SetVal(string(cacheVal))
		res, err := cacheRepo.HashGet(ctx, cacheHash, cacheKey)
		assert.NoError(t, err)
		assert.NotZero(t, res)
	})

	t.Run("success - cache client return empty string", func(t *testing.T) {
		mockedDependency.redisCmd.ExpectHGet(cacheHash, cacheKey).SetErr(redis.Nil)
		res, err := cacheRepo.HashGet(ctx, cacheHash, cacheKey)
		assert.NoError(t, err)
		assert.Zero(t, res)
	})

	t.Run("failed - cache client return error", func(t *testing.T) {
		mockedDependency.redisCmd.ExpectHGet(cacheHash, cacheKey).SetErr(errors.New("redis error"))
		res, err := cacheRepo.HashGet(ctx, cacheHash, cacheKey)
		assert.Error(t, err)
		assert.Zero(t, res)
	})
}

func TestCacheRepository_HashSet(t *testing.T) {
	mockedDependency := newMockedDependency(t)
	defer mockedDependency.close()

	ctx := mockedDependency.ctx
	cacheRepo := cacheRepo{redisClient: mockedDependency.redis}
	book := model.Book{
		ID:          int64(1),
		Title:       "Harry Potter",
		Author:      "J. K. Rowling",
		Description: "A series about wizards",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	cacheKey := "book:1"
	cacheHash := "hash:book"
	cacheVal, err := json.Marshal(book)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		mockedDependency.redisCmd.ExpectHSet(cacheHash, cacheKey, string(cacheVal)).SetVal(1)
		err := cacheRepo.HashSet(ctx, cacheHash, cacheKey, string(cacheVal))
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		mockedDependency.redisCmd.ExpectHSet(cacheHash, cacheKey, string(cacheVal)).SetErr(errors.New("redis error"))
		err := cacheRepo.Set(ctx, cacheKey, string(cacheVal))
		assert.Error(t, err)
	})
}
