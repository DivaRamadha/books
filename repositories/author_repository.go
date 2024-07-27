package repositories

import (
	"main/models"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	CreateAuthor(author *models.Author) error
	FindAuthorByID(authorID uint) (*models.Author, error)
	FindAllAuthors() ([]models.Author, error)
	UpdateAuthor(author *models.Author) error
	DeleteAuthor(id uint) error
	FindAuthorByName(name string) (*models.Author, error)
}

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{db}
}

func (r *authorRepository) CreateAuthor(author *models.Author) error {
	return r.db.Create(author).Error
}

func (r *authorRepository) FindAuthorByID(authorID uint) (*models.Author, error) {
	var author models.Author
	err := r.db.First(&author, authorID).Error
	if err != nil {
		return nil, err
	}
	return &author, err
}

func (r *authorRepository) FindAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	err := r.db.Preload("Books").Find(&authors).Error
	if err != nil {
		return nil, err
	}
	for _, author := range authors {
		author.Books = nil
	}
	return authors, err
}

func (r *authorRepository) UpdateAuthor(author *models.Author) error {
	return r.db.Save(author).Error
}

func (r *authorRepository) DeleteAuthor(id uint) error {
	return r.db.Delete(&models.Author{}, "id = ?", id).Error
}

func (r *authorRepository) FindAuthorByName(name string) (*models.Author, error) {
	var author models.Author
	err := r.db.Where("name = ?", name).First(&author).Error
	if err != nil {
		return nil, err
	}
	return &author, err
}
