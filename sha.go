package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func calculateSHA256(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

func calculateSHA512(str string) string {
	hash := sha512.Sum512([]byte(str))
	return hex.EncodeToString(hash[:])
}
