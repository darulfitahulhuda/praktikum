package routes

import (
	"main/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// Route Users
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.POST("/login", controllers.LoginUsersController)

	// Route Books
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)

	// Route Books
	e.GET("/blogs", controllers.GetBlogsController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)

	// JWT Group
	r := e.Group("/jwt")
	r.GET("/users", controllers.GetUsersController)
	r.GET("/users/:id", controllers.GetUserController)
	r.DELETE("/users/:id", controllers.DeleteUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)

	// JWT Route Books
	r.GET("/books", controllers.GetBooksController)
	r.GET("/books/:id", controllers.GetBookController)
	r.POST("/books", controllers.CreateBookController)
	r.DELETE("/books/:id", controllers.DeleteBookController)
	r.PUT("/books/:id", controllers.UpdateBookController)

	return e
}
