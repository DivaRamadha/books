package controllers

import (
	"main/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	books, err := c.bookService.FindAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}
