package utils

import (
	"crypto/rand"
	"math/big"
)

const _RANDOM = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var found = make(map[string]string)

func GenerateNewId(length int, url string) string {
	id := ""
	for {
		for i := 0; i < length; i++ {
			result, _ := rand.Int(rand.Reader, big.NewInt(62))
			id += string(_RANDOM[result.Int64()])
		}
		if _, ok := found[id]; !ok {
			break
		}
	}

	found[id] = url

	return id
}

func GetUrlFromId(id string) (string, bool) {
	val, ok := found[id]
	return val, ok
}
