package services

import (
	"fmt"
	"main/models"
	"main/repositories"
	"time"
)

type AuthorService interface {
	CreateAuthor(name string, birth string) (*models.Author, error)
	FindAuthorByID(authorID uint) (*models.Author, error)
	FindAllAuthors() ([]models.Author, error)
	UpdateAuthor(author *models.Author) (*models.Author, error)
	DeleteAuthor(id uint) error
	FindAuthorByName(name string) (*models.Author, error)
}

type authorService struct {
	authorRepo repositories.AuthorRepository
}

func NewAuthorService(authorRepo repositories.AuthorRepository) AuthorService {
	return &authorService{authorRepo}
}

func (s *authorService) CreateAuthor(name string, birth string) (*models.Author, error) {
	if name == "" || birth == "" {
		return nil, fmt.Errorf("name and birth are required")
	}

	parsedBirth, err := time.Parse("2006-01-02", birth)
	if err != nil {
		return nil, fmt.Errorf("birth should be a valid date")
	}
	if parsedBirth.After(time.Now()) {
		return nil, fmt.Errorf("birth date cannot be in the future")
	}
	author := &models.Author{
		Name:  name,
		Birth: birth,
	}

	s.authorRepo.CreateAuthor(author)

	return author, nil
}

func (s *authorService) FindAuthorByID(authorID uint) (*models.Author, error) {
	return s.authorRepo.FindAuthorByID(authorID)
}

func (s *authorService) FindAllAuthors() ([]models.Author, error) {
	return s.authorRepo.FindAllAuthors()
}

func (s *authorService) UpdateAuthor(author *models.Author) (*models.Author, error) {
	err := s.authorRepo.UpdateAuthor(author)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s *authorService) DeleteAuthor(id uint) error {
	_, err := s.FindAuthorByID(id)
	if err != nil {
		return fmt.Errorf("author not found")
	}
	return s.authorRepo.DeleteAuthor(id)
}

func (s *authorService) FindAuthorByName(name string) (*models.Author, error) {
	return s.authorRepo.FindAuthorByName(name)
}
