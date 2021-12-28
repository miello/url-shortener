package dto

import (
	"time"

	"gorm.io/gorm"
)

type URLShortener struct {
	gorm.Model
	Original string
	Shorten  string
	UrlID    string
	Expires  time.Time
}
