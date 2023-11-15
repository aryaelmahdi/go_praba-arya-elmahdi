package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"praktikum/praktikum/config"
	"praktikum/praktikum/model"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockUserModel struct{}

func (mm *MockUserModel) Register(newItem model.Users) *model.Users {
	return nil
}
func (mm *MockUserModel) GetAll() []model.Users {
	return nil
}
func (mm *MockUserModel) GetByID(id int) *model.Users {
	return nil
}
func (mm *MockUserModel) Delete(id int) {}
func (mm *MockUserModel) Update(updatedData model.Users) bool {
	return false
}
func (mm *MockUserModel) Login(email string, password string) *model.Users {
	return nil
}

func TestGetUsers(t *testing.T) {
	// Setup Controller
	var cfg = config.Config{}
	var mdl = MockUserModel{}
	var ctl = NewUserControllerInterface(&mdl, cfg)
	var e = echo.New()
	e.GET("/Users", ctl.GetAll())

	var req = httptest.NewRequest(http.MethodGet, "/Users", nil)
	var res = httptest.NewRecorder()

	e.ServeHTTP(res, req)

	type ResponseData struct {
		Data    map[string]interface{} `json:"data"`
		Message string                 `json:"message"`
	}

	var tmp = ResponseData{}

	var resData = json.NewDecoder(res.Result().Body)
	err := resData.Decode(&tmp)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Nil(t, err)
	assert.NotNil(t, tmp)
	assert.Equal(t, "success", tmp.Message)
	assert.Nil(t, tmp.Data)
}

func TestCreateUser(t *testing.T) {
	var cfg = config.Config{}
	var mdl = MockUserModel{}
	var ctl = NewUserControllerInterface(&mdl, cfg)
	var e = echo.New()
	e.POST("/Users", ctl.Register())

	t.Run("Server error", func(t *testing.T) {
		var req = httptest.NewRequest(http.MethodPost, "/Users", bytes.NewReader([]byte(`{"UserID":1, "Judul":"User 1", "Konten":"Content 1"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponseData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "something went wrong", tmp.Message)
	})

	t.Run("Success Insert", func(t *testing.T) {
		var cfg = config.Config{}
		var mdl = MockUserModel{}
		var ctl = NewUserControllerInterface(&mdl, cfg)
		var e = echo.New()
		e.POST("/Users", ctl.Register())

		var req = httptest.NewRequest(http.MethodPost, "/Users", bytes.NewReader([]byte(`{"UserID":1, "Judul":"User 1", "Konten":"Content 1"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponseData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.NotNil(t, tmp)
	})
}

func TestDeleteUser(t *testing.T) {
	var cfg = config.Config{}
	var mdl = MockUserModel{}
	var ctl = NewUserControllerInterface(&mdl, cfg)
	var e = echo.New()
	e.DELETE("/Users/:id", ctl.Delete())

	t.Run("Invalid ID", func(t *testing.T) {
		var req = httptest.NewRequest(http.MethodDelete, "/Users/satu", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)
		type ResponseData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Success Delete", func(t *testing.T) {
		var req = httptest.NewRequest(http.MethodDelete, "/Users/1", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)
		type ResponseData struct {
			Data    map[string]interface{} `json:"data"`
			Message string                 `json:"message"`
		}

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, res.Code)
	})
}
