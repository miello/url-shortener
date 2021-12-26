package utils

import (
	"crypto/rand"
	"math/big"
)

const _RANDOM = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var found = make(map[string]string)

func GenerateNewId(length int, url string) string {
	id := ""
	nBig, _ := rand.Int(rand.Reader, big.NewInt(26))
	for {
		for i := 0; i < length; i++ {
			id += string(_RANDOM[nBig.Int64()])
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
