package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type URLShortener struct {
	gorm.Model
	Original string
	Shorten  string
	UrlID    string
	UserID   uuid.UUID
	User     User
	Expires  time.Time
}
