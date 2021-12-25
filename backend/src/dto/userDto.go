package dto

import "gorm.io/gorm"

type UserDTO struct {
	gorm.Model
	Handle   string
	User     string
	Password string
}
