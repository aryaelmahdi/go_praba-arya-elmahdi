package controller

import (
	"net/http"
	"tugas/praktikum/helper"
	"tugas/praktikum/model"

	"github.com/labstack/echo/v4"
)

type TransmissionsController struct {
	model model.TransmissionModel
}

func (tc *TransmissionsController) Init(tm model.TransmissionModel) {
	tc.model = tm
}

func (tc *TransmissionsController) InsertTransmission() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := model.Transmission{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		res := tc.model.InsertTransmisson(input)
		if res == nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusCreated, helper.SetResponse("success", res))
	}
}

func (tc *TransmissionsController) GetAllTransmissions() echo.HandlerFunc {
	return func(c echo.Context) error {
		res := tc.model.GetAllTransmissions()
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (tc *TransmissionsController) GetTransmissionByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		res := tc.model.GetTransmissionsByName(name)
		if res == nil {
			c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (tc *TransmissionsController) DeleteTransmission() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		if err := tc.model.DeleteTransmission(name); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusNoContent, nil)
	}
}
