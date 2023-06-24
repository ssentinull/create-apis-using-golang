package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/ssentinull/create-apis-using-golang/internal/model"
	"github.com/ssentinull/create-apis-using-golang/internal/utils"
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
	g.PUT("/books", handler.UpdateBook)
	g.DELETE("/books/:ID", handler.DeleteBookByID)
}

func (bh *BookHTTPHandler) CreateBook(c echo.Context) error {
	input := new(model.CreateBookInput)
	if err := c.Bind(input); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	book, err := bh.BookUsecase.Create(c.Request().Context(), input.ToModel())
	if err != nil {
		logrus.Error(err)
		return c.JSON(utils.ParseHTTPErrorStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusCreated, book)
}

func (bh *BookHTTPHandler) DeleteBookByID(c echo.Context) error {
	ID, err := strconv.ParseInt(c.Param("ID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, "ID param is invalid")
	}

	err = bh.BookUsecase.DeleteByID(c.Request().Context(), ID)
	if err != nil {
		logrus.Error(err)
		return c.JSON(utils.ParseHTTPErrorStatusCode(err), err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (bh *BookHTTPHandler) FetchBooks(c echo.Context) error {
	queryParams := new(model.GetBooksQueryParams)
	if err := c.Bind(queryParams); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	books, count, err := bh.BookUsecase.FindAll(c.Request().Context(), *queryParams)
	if err != nil {
		logrus.Error(err)
		return c.JSON(utils.ParseHTTPErrorStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusOK, model.NewPaginationResponse(
		books,
		queryParams.Page,
		queryParams.Size,
		count,
	))
}

func (bh *BookHTTPHandler) FetchBookByID(c echo.Context) error {
	ID, err := strconv.ParseInt(c.Param("ID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, "ID param is invalid")
	}

	book, err := bh.BookUsecase.FindByID(c.Request().Context(), ID)
	if err != nil {
		logrus.Error(err)
		return c.JSON(utils.ParseHTTPErrorStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusOK, book)
}

func (bh *BookHTTPHandler) UpdateBook(c echo.Context) error {
	input := new(model.UpdateBookInput)
	if err := c.Bind(input); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	book, err := bh.BookUsecase.Update(c.Request().Context(), input.ToModel())
	if err != nil {
		logrus.Error(err)
		return c.JSON(utils.ParseHTTPErrorStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusOK, book)
}
