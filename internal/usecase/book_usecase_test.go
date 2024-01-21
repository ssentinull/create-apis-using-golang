package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ssentinull/create-apis-using-golang/internal/model"
	"github.com/ssentinull/create-apis-using-golang/internal/model/mock"
	"github.com/stretchr/testify/assert"
)

var (
	bookID = int64(1)
	book   = &model.Book{
		ID:          bookID,
		Title:       "Harry Potter",
		Author:      "J. K. Rowling",
		Description: "A series about wizards",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	books         = []*model.Book{book}
	lenBooks      = int64(len(books))
	findAllParams = model.GetBooksQueryParams{
		Page: 1,
		Size: 5,
	}
)

func TestBookUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedBookRepo := mock.NewMockBookRepository(ctrl)
	usecase := bookUsecase{bookRepo: mockedBookRepo}
	ctx := context.Background()

	defer func() {
		ctrl.Finish()
		ctx.Done()
	}()

	t.Run("success", func(t *testing.T) {
		mockedBookRepo.EXPECT().Create(ctx, book).Times(1).Return(nil)
		res, err := usecase.Create(ctx, book)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("failed", func(t *testing.T) {
		mockedBookRepo.EXPECT().Create(ctx, book).Times(1).Return(errors.New("db error"))
		res, err := usecase.Create(ctx, book)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestBookUsecase_DeleteByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedBookRepo := mock.NewMockBookRepository(ctrl)
	usecase := bookUsecase{bookRepo: mockedBookRepo}
	ctx := context.Background()

	defer func() {
		ctrl.Finish()
		ctx.Done()
	}()

	t.Run("success", func(t *testing.T) {
		mockedBookRepo.EXPECT().DeleteByID(ctx, bookID).Times(1).Return(nil)
		err := usecase.DeleteByID(ctx, bookID)
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		mockedBookRepo.EXPECT().DeleteByID(ctx, bookID).Times(1).Return(errors.New("db error"))
		err := usecase.DeleteByID(ctx, bookID)
		assert.Error(t, err)
	})
}

func TestBookUsecase_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedBookRepo := mock.NewMockBookRepository(ctrl)
	usecase := bookUsecase{bookRepo: mockedBookRepo}
	ctx := context.Background()

	defer func() {
		ctrl.Finish()
		ctx.Done()
	}()

	t.Run("success", func(t *testing.T) {
		mockedBookRepo.EXPECT().FindByID(ctx, bookID).Times(1).Return(book, nil)
		res, err := usecase.FindByID(ctx, bookID)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("failed", func(t *testing.T) {
		mockedBookRepo.EXPECT().FindByID(ctx, bookID).Times(1).Return(nil, errors.New("db error"))
		res, err := usecase.FindByID(ctx, bookID)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestBookUsecase_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedBookRepo := mock.NewMockBookRepository(ctrl)
	usecase := bookUsecase{bookRepo: mockedBookRepo}
	ctx := context.Background()

	defer func() {
		ctrl.Finish()
		ctx.Done()
	}()

	t.Run("success", func(t *testing.T) {
		mockedBookRepo.EXPECT().FindAll(ctx, findAllParams).Times(1).Return(books, nil)
		mockedBookRepo.EXPECT().CountAll(ctx).Times(1).Return(lenBooks, nil)

		resBooks, resCount, err := usecase.FindAll(ctx, findAllParams)
		assert.NoError(t, err)
		assert.NotNil(t, resBooks)
		assert.NotZero(t, resCount)
	})

	t.Run("failed - find all return error", func(t *testing.T) {
		mockedBookRepo.EXPECT().FindAll(ctx, findAllParams).Times(1).Return(nil, errors.New("db error"))

		resBooks, resCount, err := usecase.FindAll(ctx, findAllParams)
		assert.Error(t, err)
		assert.Nil(t, resBooks)
		assert.Zero(t, resCount)
	})

	t.Run("failed - count all return error", func(t *testing.T) {
		mockedBookRepo.EXPECT().FindAll(ctx, findAllParams).Times(1).Return(books, nil)
		mockedBookRepo.EXPECT().CountAll(ctx).Times(1).Return(int64(0), errors.New("db error"))

		resBooks, resCount, err := usecase.FindAll(ctx, findAllParams)
		assert.Error(t, err)
		assert.Nil(t, resBooks)
		assert.Zero(t, resCount)
	})
}

func TestBookUsecase_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedBookRepo := mock.NewMockBookRepository(ctrl)
	usecase := bookUsecase{bookRepo: mockedBookRepo}
	ctx := context.Background()

	defer func() {
		ctrl.Finish()
		ctx.Done()
	}()

	t.Run("success", func(t *testing.T) {
		mockedBookRepo.EXPECT().Update(ctx, book).Times(1).Return(book, nil)
		res, err := usecase.Update(ctx, book)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("failed", func(t *testing.T) {
		mockedBookRepo.EXPECT().Update(ctx, book).Times(1).Return(nil, errors.New("db error"))
		res, err := usecase.Update(ctx, book)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
