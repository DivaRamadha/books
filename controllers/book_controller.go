package controllers

import (
	"main/models"
	"main/repositories"
	"main/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	bookService services.BookService
}

func NewBookController(bookService services.BookService) *BookController {
	return &BookController{bookService}
}

func (c *BookController) CreateBook(ctx *gin.Context) {
	var input struct {
		Title    string `json:"title"`
		Isbn     string `json:"isbn"`
		AuthorId uint   `json:"author_id"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorService := services.NewAuthorService(repositories.NewAuthorRepository(ctx.MustGet("database").(*gorm.DB)))

	author, err := authorService.FindAuthorByID(input.AuthorId)
	if err != nil || author == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author_id"})
		return
	}

	book, err := c.bookService.CreateBook(input.Title, input.Isbn, input.AuthorId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) FindBookByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	book, err := c.bookService.FindBookByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) FindAllBooks(ctx *gin.Context) {
	search := ctx.Query("search")
	authorId, err := strconv.Atoi(ctx.Query("author_id"))
	if err != nil {
		authorId = 0
	}
	books, err := c.bookService.FindAllBooks(search, uint(authorId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}

// func (c *BookController) FindAllBooks(ctx *gin.Context) {
// 	books, err := c.bookService.FindAllBooks()
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, books)
// }

func (c *BookController) UpdateBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input struct {
		Title    string `json:"title"`
		Isbn     string `json:"isbn"`
		AuthorId uint   `json:"author_id"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorService := services.NewAuthorService(repositories.NewAuthorRepository(ctx.MustGet("database").(*gorm.DB)))

	author, err := authorService.FindAuthorByID(input.AuthorId)
	if err != nil || author == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author_id"})
		return
	}

	book := models.Book{ID: uint(id), Title: input.Title, Isbn: input.Isbn, AuthorId: input.AuthorId}
	updatedBook, err := c.bookService.UpdateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedBook)
}

func (c *BookController) DeleteBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	book := c.bookService.DeleteBook(uint(id))

	ctx.JSON(http.StatusOK, gin.H{"success": "deleted book", "deletedBook": book})
}
