package services

import "user-rest-app/models"
type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id string) (models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}