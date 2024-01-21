package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/ssentinull/create-apis-using-golang/internal/model"
	"github.com/ssentinull/create-apis-using-golang/internal/model/mock"
	"github.com/ssentinull/create-apis-using-golang/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestBookDeliveryHTTP_CreateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookUsecase := mock.NewMockBookUsecase(ctrl)
	httpHandler := BookHTTPHandler{BookUsecase: mockBookUsecase}
	e := echo.New()

	ID := int64(10)
	title := "Harry Potter"
	author := "J. K. Rowling"
	description := "A book about wizards"

	bookInput := model.CreateBookInput{
		Title:       title,
		Author:      author,
		Description: description,
	}

	bookInputJSON, err := json.Marshal(bookInput)
	assert.NoError(t, err)

	bookModel := model.Book{
		ID:          ID,
		Title:       title,
		Author:      author,
		Description: description,
	}

	t.Run("success", func(t *testing.T) {
		idPatch := gomonkey.ApplyFunc(utils.GenerateID, func() int64 {
			return ID
		})

		timePatch := gomonkey.ApplyFunc(time.Now, func() time.Time {
			return time.Time{}
		})

		defer func() {
			idPatch.Reset()
			timePatch.Reset()
		}()

		req := httptest.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(string(bookInputJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		mockBookUsecase.EXPECT().Create(gomock.Any(), &bookModel).Times(1).Return(&bookModel, nil)

		err := httpHandler.CreateBook(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	})

	t.Run("failed - request body is invalid", func(t *testing.T) {
		type invalidInput struct {
			Title int64 `json:"title"`
		}

		faultyInput := invalidInput{Title: 1}
		faultyInputJSON, err := json.Marshal(faultyInput)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(string(faultyInputJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		err = httpHandler.CreateBook(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("failed - create book return error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(string(bookInputJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		mockBookUsecase.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(nil, errors.New("usecase error"))

		err := httpHandler.CreateBook(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestBookDeliveryHTTP_DeleteBookByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookUsecase := mock.NewMockBookUsecase(ctrl)
	httpHandler := BookHTTPHandler{BookUsecase: mockBookUsecase}
	e := echo.New()

	ID := int64(1)

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/books", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("ID")
		ctx.SetParamValues(strconv.FormatInt(ID, 10))

		mockBookUsecase.EXPECT().DeleteByID(gomock.Any(), ID).Times(1).Return(nil)

		err := httpHandler.DeleteBookByID(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})

	t.Run("failed - id params is invalid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/books", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("ID")
		ctx.SetParamValues("invalid")

		err := httpHandler.DeleteBookByID(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("failed - find by id return error", func(t *testing.T) {
		usecaseErr := errors.New("usecase error")

		req := httptest.NewRequest(http.MethodGet, "/v1/books", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("ID")
		ctx.SetParamValues(strconv.FormatInt(ID, 10))

		mockBookUsecase.EXPECT().DeleteByID(gomock.Any(), ID).Times(1).Return(usecaseErr)

		err := httpHandler.DeleteBookByID(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), usecaseErr.Error())
	})
}

func TestBookDeliveryHTTP_FetchBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookUsecase := mock.NewMockBookUsecase(ctrl)
	httpHandler := BookHTTPHandler{BookUsecase: mockBookUsecase}
	e := echo.New()

	ID := int64(1)
	getBooksQueryParams := model.GetBooksQueryParams{
		Page: 1,
		Size: 10,
	}

	queryParams := url.Values{}
	queryParams.Add("page", fmt.Sprintf("%d", getBooksQueryParams.Page))
	queryParams.Add("size", fmt.Sprintf("%d", getBooksQueryParams.Size))
	queryParamString := queryParams.Encode()

	bookModels := []*model.Book{
		{
			ID:          ID,
			Title:       "Harry Potter",
			Author:      "J. K. Rowling",
			Description: "A book about wizards",
		},
	}

	bookModelJSON, err := json.Marshal(bookModels)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/books?%s", queryParamString), nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		mockBookUsecase.EXPECT().FindAll(gomock.Any(), getBooksQueryParams).Times(1).Return(bookModels, int64(len(bookModels)), nil)

		err := httpHandler.FetchBooks(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), string(bookModelJSON))
	})

	t.Run("failed - query param is invalid", func(t *testing.T) {
		type invalidQueryParams struct {
			Page bool `query:"page"`
		}

		faultyQueryParams := invalidQueryParams{Page: true}

		queryParams := url.Values{}
		queryParams.Add("page", strconv.FormatBool(faultyQueryParams.Page))
		queryParamString := queryParams.Encode()

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/books?%s", queryParamString), nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		err := httpHandler.FetchBooks(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("failed - find all return error", func(t *testing.T) {
		usecaseErr := errors.New("usecase error")

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/books?%s", queryParamString), nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		mockBookUsecase.EXPECT().FindAll(gomock.Any(), getBooksQueryParams).Times(1).Return(nil, int64(0), usecaseErr)

		err := httpHandler.FetchBooks(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), usecaseErr.Error())
	})
}

func TestBookDeliveryHTTP_FetchBookByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookUsecase := mock.NewMockBookUsecase(ctrl)
	httpHandler := BookHTTPHandler{BookUsecase: mockBookUsecase}
	e := echo.New()

	ID := int64(1)
	bookModel := model.Book{
		ID:          ID,
		Title:       "Harry Potter",
		Author:      "J. K. Rowling",
		Description: "A book about wizards",
	}

	bookModelJSON, err := json.Marshal(bookModel)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/books", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("ID")
		ctx.SetParamValues(strconv.FormatInt(ID, 10))

		mockBookUsecase.EXPECT().FindByID(gomock.Any(), ID).Times(1).Return(&bookModel, nil)

		err := httpHandler.FetchBookByID(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), string(bookModelJSON))
	})

	t.Run("failed - id params is invalid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/books", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("ID")
		ctx.SetParamValues("invalid")

		err := httpHandler.FetchBookByID(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "ID param is invalid")
	})

	t.Run("failed - find by id return error", func(t *testing.T) {
		usecaseErr := errors.New("usecase error")

		req := httptest.NewRequest(http.MethodGet, "/v1/books", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("ID")
		ctx.SetParamValues(strconv.FormatInt(ID, 10))

		mockBookUsecase.EXPECT().FindByID(gomock.Any(), ID).Times(1).Return(nil, usecaseErr)

		err := httpHandler.FetchBookByID(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), usecaseErr.Error())
	})
}

func TestBookDeliveryHTTP_UpdateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookUsecase := mock.NewMockBookUsecase(ctrl)
	httpHandler := BookHTTPHandler{BookUsecase: mockBookUsecase}
	e := echo.New()

	ID := int64(100)
	title := "Harry Potter"
	author := "J. K. Rowling"
	description := "A book about wizards"

	bookInput := model.UpdateBookInput{
		ID:          ID,
		Title:       title,
		Author:      author,
		Description: description,
	}

	bookInputJSON, err := json.Marshal(bookInput)
	assert.NoError(t, err)

	bookModel := model.Book{
		ID:          ID,
		Title:       title,
		Author:      author,
		Description: description,
	}

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/v1/books", strings.NewReader(string(bookInputJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		mockBookUsecase.EXPECT().Update(gomock.Any(), gomock.Any()).Times(1).Return(&bookModel, nil)

		err := httpHandler.UpdateBook(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	})

	t.Run("failed - request body is invalid", func(t *testing.T) {
		type invalidInput struct {
			Title int64 `json:"title"`
		}

		faultyInput := invalidInput{Title: 1}
		faultyInputJSON, err := json.Marshal(faultyInput)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/v1/books", strings.NewReader(string(faultyInputJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		err = httpHandler.UpdateBook(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("failed - update book return error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/v1/books", strings.NewReader(string(bookInputJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		mockBookUsecase.EXPECT().Update(gomock.Any(), gomock.Any()).Times(1).Return(nil, errors.New("usecase error"))

		err := httpHandler.UpdateBook(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
