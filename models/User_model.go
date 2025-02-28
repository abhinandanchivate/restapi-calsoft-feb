package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string    `validate:"required,min=3,max=50"`
	Email    string    `gorm:"unique" validate:"required,email"`
	Password string    `validate:"required,min=6,max=50"`
	RoleID   uuid.UUID `validate:"required"`
	Role     Role      `gorm:"foreignKey:RoleID"`
}

/*
RoleID: A required UUID field that acts as a foreign key pointing to the Role model.
Role: A reference to the Role model, establishing a one-to-many relationship (each user has one role).
*/


