package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id" gorm:"primaryKey;type:int(10) AUTO_INCREMENT"`
	Name     string `json:"name" form:"name" gorm:"type:varchar(20)"`
	Email    string `json:"email" form:"email" gorm:"type:varchar(50)"`
	Password string `json:"password" form:"password" gorm:"type:varchar(20)"`
}

var users []User

// -------------------- controller --------------------

func setResponse(message string, data any) map[string]any {
	if data == nil {
		return map[string]any{
			"message": message,
		}
	}
	return map[string]any{
		"message": message,
		"data":    data,
	}
}

// get all users
func GetUsersController(c echo.Context) error {
	//edited
	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	var path = c.Param("id")
	id, err := strconv.Atoi(path)

	if err != nil {
		return c.JSON(http.StatusBadRequest, setResponse(err.Error(), nil))
	}

	if id < 0 || id > len(users)-1 {
		return c.JSON(http.StatusNotFound, setResponse("data not found", nil))
	}

	userData := users[id]
	return c.JSON(http.StatusOK, setResponse("success", userData))
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	var path = c.Param("id")
	id, err := strconv.Atoi(path)
	if err != nil {
		return c.JSON(http.StatusBadRequest, setResponse(err.Error(), nil))
	}

	if id < 0 || id > len(users)-1 {
		return c.JSON(http.StatusNotFound, setResponse("data not found", nil))
	}

	users[id] = User{}
	return c.JSON(http.StatusNoContent, nil)
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	var path = c.Param("id")
	var updatedData = User{}
	c.Bind(&updatedData)
	id, err := strconv.Atoi(path)
	if err != nil {
		return c.JSON(http.StatusBadRequest, setResponse(err.Error(), nil))
	}

	if id < 0 || id > len(users)-1 {
		return c.JSON(http.StatusNotFound, setResponse("data not found", nil))
	}

	users[id].Name = updatedData.Name
	users[id].Password = updatedData.Password
	users[id].Email = updatedData.Email
	return c.JSON(http.StatusOK, setResponse("success", nil))
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
