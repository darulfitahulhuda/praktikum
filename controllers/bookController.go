package controllers

import (
	"main/lib/database"
	"main/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all Books
func GetBooksController(c echo.Context) error {
	result, e := database.GetBooks()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"data":    result,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, e := database.GetBookById(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"data":    result,
	})

}

// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	result, e := database.CreateBook(book)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"data":    result,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	_, e := database.DeleteBook(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book := models.Book{}
	c.Bind(&book)

	_, e := database.UpdateBook(id, book)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
	})
}
