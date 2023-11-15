package controller

import (
	"net/http"
	"tugas/praktikum/helper"
	"tugas/praktikum/model"

	"github.com/labstack/echo/v4"
)

type AspirationsController struct {
	model model.AspirationsModel
}

func (ac *AspirationsController) Init(am model.AspirationsModel) {
	ac.model = am
}

func (ac *AspirationsController) InsertAspiration() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := model.Aspirations{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}
		res, err := ac.model.InsertAspiration(input)

		if res == nil || err != nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}

		return c.JSON(http.StatusCreated, helper.SetResponse("success", res))
	}
}

func (a *AspirationsController) DeleteAspiration() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if err := a.model.DeleteAspiration(id); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("cannot delete data", nil))
		}
		return c.JSON(http.StatusNoContent, nil)
	}
}

func (a *AspirationsController) GetAllAspirations() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := a.model.GetAllAspirations()
		if res == nil || err != nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (a *AspirationsController) GetAspirationByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		res := a.model.GetAspirationByID(id)
		if res == nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid id", nil))
		}
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}
