package repositories

import (
	"main/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	FindBookByID(bookID uint) (*models.Book, error)
	FindAllBooks() ([]models.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) CreateBook(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) FindBookByID(bookID uint) (*models.Book, error) {
	var book models.Book
	err := r.db.Preload("Author").First(&book, bookID).Error
	if err != nil {
		return nil, err
	}
	return &book, err
}

func (r *bookRepository) FindAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Preload("Author").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, err
}
