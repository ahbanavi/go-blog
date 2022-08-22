package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `json:"email" gorm:"type:varchar(100);unique_index" validate:"required,email"`
	PasswordHash string `json:"-" gorm:"type:char(60)"`
}
