package controllers

import (
	"main/models"
	"main/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	authorService services.AuthorService
}

func NewAuthorController(authorService services.AuthorService) *AuthorController {
	return &AuthorController{authorService}
}

// func NewRouter(bookController *BookController) *gin.Engine {
// 	router := gin.Default()

// 	// Define routes
// 	router.POST("/author", bookController.CreateBook)
// 	router.GET("/author/:id", bookController.FindBookByID)
// 	router.GET("/author", bookController.FindAllBooks)

// 	return router
// }

func (c *AuthorController) CreateAuthor(ctx *gin.Context) {
	var author models.Author
	if err := ctx.ShouldBindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdAuthor, err := c.authorService.CreateAuthor(author.Name, author.Birth)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, createdAuthor)
}

func (c *AuthorController) FindAllAuthors(ctx *gin.Context) {
	authors, err := c.authorService.FindAllAuthors()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, authors)
}

func (c *AuthorController) FindAuthorByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	author, err := c.authorService.FindAuthorByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Author not found"})
		return
	}
	ctx.JSON(http.StatusOK, author)
}
