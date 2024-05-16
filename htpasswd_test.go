package main

import (
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestGenerateHtpasswdEntry(t *testing.T) {
	username := "testuser"
	password := "testpassword"

	entry, err := generateHtpasswdEntry(username, password)
	if err != nil {
		t.Fatalf("Error generating htpasswd entry: %v", err)
	}

	if !strings.HasPrefix(entry, username+":") {
		t.Errorf("Entry does not contain username: %s", entry)
	}

	parts := strings.SplitN(entry, ":", 2)
	if len(parts) != 2 {
		t.Fatalf("Invalid htpasswd entry format: %s", entry)
	}
	hash := parts[1]

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Errorf("Password does not match hash: %v", err)
	}
}
