package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type URLShortener struct {
	gorm.Model
	Original string
	UrlID    string
	UserID   uuid.UUID `gorm:"default:null"`
	User     User
	Expires  time.Time
}
