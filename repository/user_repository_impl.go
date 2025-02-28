package repository

import (
	"user-rest-app/config"
	"user-rest-app/models"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// NewUserRepositoryImpl creates a new instance of UserRepositoryImpl
func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{DB: config.ConnectDB()}
}

// CreateUser inserts a new user record
func (repo *UserRepositoryImpl) CreateUser(user *models.User) error {
	return repo.DB.Create(user).Error
}

// GetUserByID retrieves a user by ID, including their role
func (repo *UserRepositoryImpl) GetUserByID(id string) (models.User, error) {
	var user models.User
	err := repo.DB.Preload("Role").Where("id = ?", id).First(&user).Error
	return user, err
}

// GetAllUsers retrieves all users, including their roles
func (repo *UserRepositoryImpl) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := repo.DB.Preload("Role").Find(&users).Error
	return users, err
}

// UpdateUser updates an existing user record
func (repo *UserRepositoryImpl) UpdateUser(user *models.User) error {
	return repo.DB.Save(user).Error
}

// DeleteUser removes a user record by ID
func (repo *UserRepositoryImpl) DeleteUser(id string) error {
	return repo.DB.Delete(&models.User{}, "id = ?", id).Error
}