package controllers

import (
	"main/models"
	"main/repositories"
	"main/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthorController struct {
	authorService services.AuthorService
}

func NewAuthorController(authorService services.AuthorService) *AuthorController {
	return &AuthorController{authorService}
}
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

func (c *AuthorController) UpdateAuthor(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var author models.Author
	if err := ctx.ShouldBindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if author.Name == "" && author.Birth == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name and Birth are required"})
		return
	}

	parsedBirth, err := time.Parse("2006-01-02", author.Birth)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Birth should be a valid date"})
		return
	}
	if parsedBirth.After(time.Now()) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Birth date cannot be in the future"})
		return
	}

	foundAuthor, err := c.authorService.FindAuthorByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Author not found"})
		return
	}
	if foundAuthor != nil && foundAuthor.ID != uint(id) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Author not found"})
		return
	}

	author.ID = uint(id)
	updatedAuthor, err := c.authorService.UpdateAuthor(&author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedAuthor)
}

func (c *AuthorController) DeleteAuthor(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	bookService := services.NewBookService(repositories.NewBookRepository(ctx.MustGet("database").(*gorm.DB)))
	books, err := bookService.FindAllBooks("", uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, book := range books {
		bookService.DeleteBook(book.ID)
	}
	author := c.authorService.DeleteAuthor(uint(id))
	ctx.JSON(http.StatusOK, gin.H{"success": "deleted author", "author": author})
}

func (c *AuthorController) FindAuthorByName(ctx *gin.Context) {
	name := ctx.Param("name")
	author, err := c.authorService.FindAuthorByName(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Author not found"})
		return
	}
	ctx.JSON(http.StatusOK, author)
}
