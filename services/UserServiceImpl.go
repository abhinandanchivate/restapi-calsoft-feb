package services

import (
	"user-rest-app/models"
	"user-rest-app/repository"
)
type UserServiceImpl struct {
	UserRepo repository.UserRepository
}

// NewUserServiceImpl creates a new instance of UserServiceImpl
func NewUserServiceImpl(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepo: userRepo}
}

// CreateUser handles creating a new user
func (service *UserServiceImpl) CreateUser(user *models.User) error {
	return service.UserRepo.CreateUser(user)
}

// GetUserByID retrieves a user by ID
func (service *UserServiceImpl) GetUserByID(id string) (models.User, error) {
	return service.UserRepo.GetUserByID(id)
}

// GetAllUsers retrieves all users
func (service *UserServiceImpl) GetAllUsers() ([]models.User, error) {
	return service.UserRepo.GetAllUsers()
}

// UpdateUser updates an existing user
func (service *UserServiceImpl) UpdateUser(user *models.User) error {
	return service.UserRepo.UpdateUser(user)
}

// DeleteUser deletes a user by ID
func (service *UserServiceImpl) DeleteUser(id string) error {
	return service.UserRepo.DeleteUser(id)
}
