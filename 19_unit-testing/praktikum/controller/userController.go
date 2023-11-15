package controller

import (
	"fmt"
	"net/http"
	"praktikum/praktikum/config"
	"praktikum/praktikum/helper"
	"praktikum/praktikum/model"
	"strconv"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	model  model.UserInterface
	config config.Config
}

type UserInterface interface {
	Register() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
	Login() echo.HandlerFunc
}

func NewUserControllerInterface(model model.UserInterface, config config.Config) UserController {
	return UserController{model: model, config: config}
}

// func (uc *UserController) Init(um model.UserModel, cfg config.Config) {
// 	uc.model = um
// 	uc.config = cfg
// }

func (uc *UserController) Register() echo.HandlerFunc {
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
		var token = c.Get("user")
		var jwtClaims = helper.ExtractToken(token.(*jwt.Token))
		id, err := strconv.Atoi(param)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid input", nil))
		}

		if jwtClaims == nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("no token", nil))
		}

		res := uc.model.GetByID(id)

		if res == nil {
			return c.JSON(http.StatusNotFound, helper.SetResponse("user not found", nil))
		}

		return c.JSON(http.StatusOK, helper.SetResponse("success", res))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = model.Login{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.SetResponse("invalid user input", nil))
		}

		res := uc.model.Login(input.Email, input.Password)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("something went wrong", nil))
		}

		if res.ID == 0 {
			return c.JSON(http.StatusNotFound, helper.SetResponse("data not found", nil))
		}

		jwtToken := helper.GenerateJWT(uc.config.Secret, res.Name, fmt.Sprint(res.ID))

		if jwtToken == nil {
			return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data", nil))
		}

		jwtToken["info"] = res

		return c.JSON(http.StatusOK, helper.SetResponse("success", jwtToken))
	}
}
