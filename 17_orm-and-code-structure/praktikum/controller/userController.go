package controller

import (
	"net/http"
	"praktikum/praktikum/helper"
	"praktikum/praktikum/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	model model.UserModel
}

func (uc *UserController) Init(um model.UserModel) {
	uc.model = um
}

func (uc *UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := model.Users{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}

		var res = uc.model.Register(input)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}

		return c.JSON(http.StatusCreated, helper.SetResponse("success", res))
	}
}

func (uc *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		input := model.Users{}

		id, err := strconv.Atoi(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("failed to update data", nil))
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid user input", nil))
		}
		input.ID = uint(id)
		res := uc.model.Update(input)

		if !res {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot update data", nil))
		}

		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (uc *UserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res := uc.model.GetAll()
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (uc *UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		id, err := strconv.Atoi(param)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid id", nil))
		}

		uc.model.Delete(id)

		return c.JSON(http.StatusNoContent, nil)
	}
}

func (uc *UserController) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var param = c.Param("id")
		id, err := strconv.Atoi(param)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}

		res := uc.model.GetByID(id)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}

		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}
