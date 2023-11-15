package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"praktikum/praktikum/model"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockBookModel struct{}

func (mm *MockBookModel) Insert(newItem model.Books) *model.Books {
	return nil
}
func (mm *MockBookModel) GetAllBooks() []model.Books {
	return nil
}
func (mm *MockBookModel) GetBookByID(id int) *model.Books {
	return nil
}
func (mm *MockBookModel) DeleteBook(id int) {}
func (mm *MockBookModel) UpdateBook(updatedData model.Books) bool {
	return false
}

type SuccessBookModel struct{}

func (smm *SuccessMockModel) Insert(newItem model.Books) *model.Books {
	return &newItem
}
func (smm *SuccessMockModel) GetAllBooks() []model.Books {
	return nil
}
func (smm *SuccessMockModel) DeleteBook(id int) {}
func (smm *SuccessMockModel) GetBookByID(id int) *model.Books {
	return nil
}
func (smm *SuccessMockModel) UpdateBook(updatedData model.Books) bool {
	return false
}

func TestGetBooks(t *testing.T) {
	// Setup Controller
	var mdl = MockBookModel{}
	var ctl = NewBookControllerInterface(&mdl)
	var e = echo.New()
	e.GET("/Books", ctl.GetAllBooks())

	var req = httptest.NewRequest(http.MethodGet, "/Books", nil)
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

func TestCreateBook(t *testing.T) {
	var mdl = MockBookModel{}
	var ctl = NewBookControllerInterface(&mdl)
	var e = echo.New()
	e.POST("/Books", ctl.Create())

	t.Run("Server error", func(t *testing.T) {
		var req = httptest.NewRequest(http.MethodPost, "/Books", bytes.NewReader([]byte(`{"UserID":1, "Judul":"Book 1", "Konten":"Content 1"}`)))
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
		var mdl = SuccessMockModel{}
		var ctl = NewBookControllerInterface(&mdl)
		var e = echo.New()
		e.POST("/Books", ctl.Create())

		var req = httptest.NewRequest(http.MethodPost, "/Books", bytes.NewReader([]byte(`{"UserID":1, "Judul":"Book 1", "Konten":"Content 1"}`)))
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

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "success", tmp.Message)
		assert.NotNil(t, tmp)
	})
}

func TestDeleteBook(t *testing.T) {
	var mdl = MockBookModel{}
	var ctl = NewBookControllerInterface(&mdl)
	var e = echo.New()
	e.DELETE("/Books/:id", ctl.DeleteBook())

	t.Run("Invalid ID", func(t *testing.T) {
		var req = httptest.NewRequest(http.MethodDelete, "/Books/satu", nil)
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
		var req = httptest.NewRequest(http.MethodDelete, "/Books/1", nil)
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
