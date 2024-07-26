package repositories

import (
	"main/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByID(userID uint) (*models.User, error)
	FindAllUsers() ([]models.User, error)
	CreateUser(user *models.User) error
	FindUserByUserName(username string) (*models.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) FindUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *userRepository) FindAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, err
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindUserByUserName(username string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
