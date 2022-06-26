package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/create-apis-using-golang/pkg/model"
)

type BookHTTPHandler struct {
	BookUsecase model.BookUsecase
}

func NewBookHTTPHandler(e *echo.Echo, bu model.BookUsecase) {
	handler := BookHTTPHandler{BookUsecase: bu}

	g := e.Group("/v1")
	g.POST("/books", handler.CreateBook)
	g.GET("/books", handler.FetchBooks)
	g.GET("/books/:ID", handler.FetchBookByID)
}

func (bh *BookHTTPHandler) CreateBook(c echo.Context) error {
	input := new(model.CreateBookInput)
	if err := c.Bind(input); err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	book, err := bh.BookUsecase.CreateBook(c.Request().Context(), input)
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, book)
}

func (bh *BookHTTPHandler) FetchBooks(c echo.Context) error {
	books, err := bh.BookUsecase.GetBooks(c.Request().Context())
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, books)
}

func (bh *BookHTTPHandler) FetchBookByID(c echo.Context) error {
	ID, err := strconv.ParseInt(c.Param("ID"), 10, 64)
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, "url param is faulty")
	}

	book, err := bh.BookUsecase.GetBookByID(c.Request().Context(), ID)
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, book)
}
