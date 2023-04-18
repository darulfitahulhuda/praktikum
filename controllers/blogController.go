package controllers

import (
	"main/lib/database"
	"main/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all Blogs
func GetBlogsController(c echo.Context) error {
	result, e := database.GetBlogs()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"data":    result,
	})
}

// get blog by id
func GetBlogController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, e := database.GetBlogById(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog",
		"data":    result,
	})

}

// create new blog
func CreateBlogController(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)

	result, e := database.CreateBlog(blog)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"data":    result,
	})
}

// delete blog by id
func DeleteBlogController(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	_, e := database.DeleteBlog(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
	})
}

// update blog by id
func UpdateBlogController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	blog := models.Blog{}
	c.Bind(&blog)

	_, e := database.UpdateBlog(id, blog)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
	})
}
