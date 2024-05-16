package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func generateHtpasswdEntry(username, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", username, hash), nil
}
