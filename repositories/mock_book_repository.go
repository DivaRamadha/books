package repositories

import (
	"main/models"

	"github.com/stretchr/testify/mock"
)

// MockBookRepository is a mock implementation of BookRepository
type MockBookRepository struct {
	mock.Mock
}

// CreateBook mocks the creation of a book
func (m *MockBookRepository) CreateBook(book *models.Book) error {
	args := m.Called(book)
	return args.Error(0)
}
