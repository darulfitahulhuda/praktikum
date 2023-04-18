package controllers

import (
	"main/lib/database"
	"main/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	result, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"data":    result,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	result, e := database.GetUserById(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"data":    result,
	})

}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result, e := database.CreateUser(user)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"data":    result,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	_, e := database.DeleteUser(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := models.User{}
	c.Bind(&user)

	_, e := database.UpdateUser(id, user)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}

// func login user
func LoginUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result, e := database.LoginUsers(&user)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"data":    result,
	})
}
