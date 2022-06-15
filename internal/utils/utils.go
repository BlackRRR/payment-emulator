package utils

import "math/rand"

const (
	AvailableSymbolInUUID = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"
	UUIDLength            = 7
)

func GetUUID() string {
	var key string

	rs := []rune(AvailableSymbolInUUID)
	lenOfArray := len(rs)

	for i := 0; i < UUIDLength; i++ {
		key += string(rs[rand.Intn(lenOfArray)])
	}
	return key
}
