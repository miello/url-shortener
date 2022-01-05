package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type URLShortener struct {
	gorm.Model
	Original string
	UrlID    string
	UserID   uuid.UUID `gorm:"default:null"`
	User     User
	Expires  string
}
