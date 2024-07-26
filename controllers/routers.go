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

		// Author routes
		api.GET("/authors", authorController.FindAllAuthors)
		api.GET("/authors/:id", authorController.FindAuthorByID)
		api.POST("/authors", authorController.CreateAuthor)

		// User routes
		api.GET("/users/:id", userController.FindUserByID)
		api.GET("/users", userController.FindAllUsers)
	}
	return router
}
