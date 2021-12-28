package utils

import (
	"crypto/rand"
	"math/big"

	"github.com/miello/url-shortener/backend/src/dto"
)

const _RANDOM = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateNewId(length int, url string) string {
	id := ""
	db := GetDB()
	for {
		id = ""
		for i := 0; i < length; i++ {
			result, _ := rand.Int(rand.Reader, big.NewInt(62))
			id += string(_RANDOM[result.Int64()])
		}

		var cnt int64

		db.Where(&dto.URLShortener{UrlID: id}).Count(&cnt)
		if cnt == 0 {
			break
		}
	}

	return id
}

func GetUrlFromId(id string) (string, bool) {
	var fetchUrl dto.URLShortener
	db.Where(&dto.URLShortener{UrlID: id}).First(&fetchUrl)
	return fetchUrl.Original, fetchUrl.Original != ""
}
