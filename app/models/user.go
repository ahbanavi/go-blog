package models

import "gorm.io/gorm"

type User struct {
	// TODO: add json tags
	gorm.Model
	Email        string
	PasswordHash string
}
