package models

import (
	"github.com/google/uuid"
)

type Role struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string    `gorm:"unique;not null" validate:"required,min=3,max=20"`
}

