package dto

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Handle   string
	User     string
	Password string
}
