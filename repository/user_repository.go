package repository

import "user-rest-app/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id string) (models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}
