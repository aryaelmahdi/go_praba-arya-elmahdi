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

type MockBlogModel struct{}

func (mm *MockBlogModel) CreateBlog(newItem model.Blog) *model.Blog {
	return nil
}
func (mm *MockBlogModel) GetAllBlogs() []model.Blog {
	return nil
}
func (mm *MockBlogModel) GetBlogByID(id int) *model.Blog {
	return nil
}
func (mm *MockBlogModel) DeleteBlog(id int) {}
func (mm *MockBlogModel) UpdateBlog(updatedData model.Blog) bool {
	return false
}

type SuccessMockModel struct{}

func (smm *SuccessMockModel) CreateBlog(newItem model.Blog) *model.Blog {
	return &newItem
}
func (smm *SuccessMockModel) GetAllBlogs() []model.Blog {
	return nil
}
func (smm *SuccessMockModel) DeleteBlog(id int) {}
func (smm *SuccessMockModel) GetBlogByID(id int) *model.Blog {
	return nil
}
func (smm *SuccessMockModel) UpdateBlog(updatedData model.Blog) bool {
	return false
}

func TestGetBlogs(t *testing.T) {
	// Setup Controller
	var mdl = MockBlogModel{}
	var ctl = NewBlogControllerInterface(&mdl)
	var e = echo.New()
	e.GET("/blogs", ctl.GetAllBlogs())

	var req = httptest.NewRequest(http.MethodGet, "/blogs", nil)
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

func TestCreateBlog(t *testing.T) {
	var mdl = MockBlogModel{}
	var ctl = NewBlogControllerInterface(&mdl)
	var e = echo.New()
	e.POST("/blogs", ctl.CreateBlog())

	t.Run("Server error", func(t *testing.T) {
		var req = httptest.NewRequest(http.MethodPost, "/blogs", bytes.NewReader([]byte(`{"UserID":1, "Judul":"Blog 1", "Konten":"Content 1"}`)))
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
		var ctl = NewBlogControllerInterface(&mdl)
		var e = echo.New()
		e.POST("/blogs", ctl.CreateBlog())

		var req = httptest.NewRequest(http.MethodPost, "/blogs", bytes.NewReader([]byte(`{"UserID":1, "Judul":"Blog 1", "Konten":"Content 1"}`)))
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

func TestDeleteBlog(t *testing.T) {
	var mdl = MockBlogModel{}
	var ctl = NewBlogControllerInterface(&mdl)
	var e = echo.New()
	e.DELETE("/blogs/:id", ctl.DeleteBlog())

	t.Run("Invalid ID", func(t *testing.T) {
		var req = httptest.NewRequest(http.MethodDelete, "/blogs/satu", nil)
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
		var req = httptest.NewRequest(http.MethodDelete, "/blogs/1", nil)
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
