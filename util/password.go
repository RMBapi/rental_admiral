package util

import (
	"crypto/rand"
	"encoding/hex"
)

func GeneratePassword(length int) (string, error) {
	bytes := make([]byte, length/2)
	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	hexID := hex.EncodeToString(bytes)

	return hexID, nil
}


func UserPassword() (string, error) {
	keyLength := 6
	key, err := GeneratePassword(keyLength)
	if err != nil {
		return "", err
	}

	Hashpassword, err := Hashpassword(key)

	return Hashpassword, nil

}

