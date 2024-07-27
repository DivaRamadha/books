package controllers

import (
	"main/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(bookController *BookController, authorController *AuthorController, userController *UserController, authController *AuthController, db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.DBMiddleware(db))

	router.POST("/auth/login", authController.Login)
	router.POST("/auth/register", authController.Register)

	api := router.Group("/")
	api.Use(middlewares.AuthMiddleware(db))
	{
		// Book routes
		api.POST("/books", bookController.CreateBook)
		api.GET("/books/:id", bookController.FindBookByID)
		api.GET("/books", bookController.FindAllBooks)
		router.PUT("/books/:id", bookController.UpdateBook)
		api.DELETE("/books/:id", bookController.DeleteBook)

		// Author routes
		api.GET("/authors", authorController.FindAllAuthors)
		api.GET("/authors/:id", authorController.FindAuthorByID)
		api.PUT("/authors/:id", authorController.UpdateAuthor)
		api.DELETE("/authors/:id", authorController.DeleteAuthor)
		api.GET("/authors/name/:name", authorController.FindAuthorByName)
		api.POST("/authors", authorController.CreateAuthor)

		// User routes
		api.GET("/users/:id", userController.FindUserByID)
		api.GET("/users", userController.FindAllUsers)
	}
	return router
}
