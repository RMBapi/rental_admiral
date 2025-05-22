package util

import (
	"crypto/rand"
	"encoding/hex"

	"example.com/rental/db"
)

func GenerateHexID(length int) (string, error) {
	bytes := make([]byte, length/2)
	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	hexID := hex.EncodeToString(bytes)

	return hexID, nil
}

func UniqueCheck(key string) (bool, error) {
	count,err := db.UniqueHashIdCheck(key)
    return count > 0, err
}

func UserHashId() (string, error) {
	keyLength := 4

	for {
		key, err := GenerateHexID(keyLength)
		if err != nil {
			return "", err
		}

		is_exsists, err := UniqueCheck(key)

		if err != nil {
			return "", err
		}

		if is_exsists == false {
			return key, nil
		}
	}
}
