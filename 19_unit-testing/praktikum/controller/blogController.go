package controller

import (
	"net/http"
	"praktikum/praktikum/helper"
	"praktikum/praktikum/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BlogController struct {
	model model.BlogInterface
}

type BlogControllerTest struct {
	model model.BlogInterface
}

type BlogControllerInterface interface {
	CreateBlog() echo.HandlerFunc
	UpdateBlog() echo.HandlerFunc
	GetAllBlogs() echo.HandlerFunc
	GetBlogById() echo.HandlerFunc
	DeleteBlog() echo.HandlerFunc
}

func NewBlogControllerInterface(model model.BlogInterface) BlogControllerInterface {
	return &BlogController{model: model}
}

// func (bc *BlogController) Init(bm model.BlogModel) {
// 	bc.model = bm
// }

func (bc *BlogController) CreateBlog() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := model.Blog{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}

		var res = bc.model.CreateBlog(input)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}

		return c.JSON(http.StatusCreated, helper.SetResponse("success", res))
	}
}

func (bc *BlogController) UpdateBlog() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		input := model.Blog{}

		id, err := strconv.Atoi(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("failed to update data", nil))
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid user input", nil))
		}
		input.ID = uint(id)
		res := bc.model.UpdateBlog(input)

		if !res {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot update data", nil))
		}

		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (bc *BlogController) GetAllBlogs() echo.HandlerFunc {
	return func(c echo.Context) error {
		res := bc.model.GetAllBlogs()
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (bc *BlogController) DeleteBlog() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		id, err := strconv.Atoi(param)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid id", nil))
		}

		bc.model.DeleteBlog(id)

		return c.JSON(http.StatusNoContent, nil)
	}
}

func (bc *BlogController) GetBlogById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var param = c.Param("id")
		id, err := strconv.Atoi(param)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}

		res := bc.model.GetBlogByID(id)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}

		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}
