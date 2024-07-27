package repositories

import (
	"main/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	FindBookByID(bookID uint) (*models.Book, error)
	FindAllBooks(search string, authorID uint) ([]models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
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

func (r *bookRepository) FindAllBooks(search string, authorID uint) ([]models.Book, error) {
	var books []models.Book
	if search != "" && authorID == 0 {
		err := r.db.Preload("Author").Where("title ILIKE ?", "%"+search+"%").Find(&books).Error
		if err != nil {
			return nil, err
		}
	} else if search == "" && authorID != 0 {
		err := r.db.Preload("Author").Where("author_id = ?", authorID).Find(&books).Error
		if err != nil {
			return nil, err
		}
	} else if search != "" && authorID != 0 {
		err := r.db.Preload("Author").Where("title LIKE ? AND author_id = ?", "%"+search+"%", authorID).Find(&books).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := r.db.Preload("Author").Find(&books).Error
		if err != nil {
			return nil, err
		}
	}
	return books, nil
}

func (r *bookRepository) UpdateBook(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) DeleteBook(id uint) error {
	return r.db.Delete(&models.Book{}, "id = ?", id).Error
}
