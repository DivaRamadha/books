package services

import (
	"main/models"
	"main/repositories"
)

type UserService interface {
	FindUserByID(userID uint) (*models.User, error)
	FindAllUsers() ([]models.User, error)
	CreateUser(user *models.User) error
	FindUserByUserName(username string) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) FindUserByID(userID uint) (*models.User, error) {
	return s.userRepo.FindUserByID(userID)
}
func (s *userService) FindAllUsers() ([]models.User, error) {
	return s.userRepo.FindAllUsers()
}
func (s *userService) CreateUser(user *models.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *userService) FindUserByUserName(username string) (*models.User, error) {
	return s.userRepo.FindUserByUserName(username)
}
