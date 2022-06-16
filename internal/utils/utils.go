package utils

import (
	"math/rand"
	"time"
)

const (
	AvailableSymbolInUUID = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"
	UUIDLength            = 7
)

func GetUUID() string {
	rand.Seed(time.Now().UnixNano())
	var key string

	rs := []rune(AvailableSymbolInUUID)
	lenOfArray := len(rs)

	for i := 0; i < UUIDLength; i++ {
		key += string(rs[rand.Intn(lenOfArray)])
	}
	return key
}
