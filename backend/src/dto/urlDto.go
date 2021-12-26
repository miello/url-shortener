package dto

import (
	"time"

	"gorm.io/gorm"
)

type URLShortener struct {
	gorm.Model
	Original string
	Shorten  string
	UserID   uint
	Expires  time.Time
	User     User
}
