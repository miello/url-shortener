package dto

import "gorm.io/gorm"

type URLShortenerDTO struct {
	gorm.Model
	Original string
	Shorten  string
	UserID   uint
	User     UserDTO
}
