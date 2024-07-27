package services

import (
	"main/models"
	"main/repositories"
)

type BookService interface {
	CreateBook(title, isbn string, authorId uint) (*models.Book, error)
	FindBookByID(bookID uint) (*models.Book, error)
	FindAllBooks(search string, authorID uint) ([]models.Book, error)
	UpdateBook(book *models.Book) (*models.Book, error)
	DeleteBook(id uint) error
}

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) BookService {
	return &bookService{bookRepo}
}

func (s *bookService) CreateBook(title, isbn string, authorId uint) (*models.Book, error) {
	book := models.Book{Title: title, Isbn: isbn, AuthorId: authorId}
	if err := s.bookRepo.CreateBook(&book); err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *bookService) FindBookByID(bookID uint) (*models.Book, error) {
	book, err := s.bookRepo.FindBookByID(bookID)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *bookService) FindAllBooks(search string, authorID uint) ([]models.Book, error) {
	books, err := s.bookRepo.FindAllBooks(search, authorID)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *bookService) UpdateBook(book *models.Book) (*models.Book, error) {
	err := s.bookRepo.UpdateBook(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *bookService) DeleteBook(id uint) error {
	_, err := s.bookRepo.FindBookByID(id)
	if err != nil {
		return err
	}
	return s.bookRepo.DeleteBook(id)
}
