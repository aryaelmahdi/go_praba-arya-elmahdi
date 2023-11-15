package controller

import (
	"net/http"
	"praktikum/praktikum/helper"
	"praktikum/praktikum/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	model model.BookInterface
}

// func (bc *BookController) Init(bm model.BookModel) {
// 	bc.model = bm
// }

type BookControllerInterface interface {
	Create() echo.HandlerFunc
	GetAllBooks() echo.HandlerFunc
	GetBookByID() echo.HandlerFunc
	DeleteBook() echo.HandlerFunc
	UpdateBook() echo.HandlerFunc
}

func NewBookControllerInterface(model model.BookInterface) *BookController {
	return &BookController{model: model}
}

func (bc *BookController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := model.Books{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}

		var res = bc.model.Insert(input)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}

		return c.JSON(http.StatusCreated, helper.SetResponse("success", res))
	}
}

func (bc *BookController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		input := model.Books{}

		id, err := strconv.Atoi(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("failed to update data", nil))
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid user input", nil))
		}
		input.ID = uint(id)
		res := bc.model.UpdateBook(input)

		if !res {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot update data", nil))
		}

		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (bc *BookController) GetAllBooks() echo.HandlerFunc {
	return func(c echo.Context) error {
		res := bc.model.GetAllBooks()
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (bc *BookController) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		id, err := strconv.Atoi(param)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid id", nil))
		}

		bc.model.DeleteBook(id)

		return c.JSON(http.StatusNoContent, nil)
	}
}

func (bc *BookController) GetBookByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var param = c.Param("id")
		id, err := strconv.Atoi(param)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}

		res := bc.model.GetBookByID(id)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}

		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}
