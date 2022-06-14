package utils

import "math/rand"

const (
	AvailableSymbolInHash = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"
	HashLength            = 7
)

func GetHash() string {
	var key string

	rs := []rune(AvailableSymbolInHash)
	lenOfArray := len(rs)

	for i := 0; i < HashLength; i++ {
		key += string(rs[rand.Intn(lenOfArray)])
	}
	return key
}
