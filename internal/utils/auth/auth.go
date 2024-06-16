package auth

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) (string, error) {
	hashedPassword := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hashedPassword[:]), nil
}

func VerifyPassword(password string, hashedPassword string) bool {
	userHashedPassword := sha256.Sum256([]byte(password))
	return hex.EncodeToString(userHashedPassword[:]) == hashedPassword
}
