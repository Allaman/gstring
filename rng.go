package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func pseudoRandomGenerator(start, end, n int, delimiter string) (string, error) {
	if start > end {
		return "", fmt.Errorf("start should be less than or equal to end")
	}
	if n <= 0 {
		return "", fmt.Errorf("number of random numbers to generate should be greater than zero")
	}

	rngs := make([]int, n)
	for i := 0; i < n; i++ {
		rngs[i] = rand.Intn(end-start+1) + start
	}
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(rngs)), delimiter), "[]"), nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const defaultSpecialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?"

func generateRandomPasswords(length, count, minSpecial int, specialChars string) ([]string, error) {
	cs := charset
	if specialChars != "" {
		cs += specialChars
	}
	if minSpecial > 0 && specialChars == "" {
		return nil, fmt.Errorf("--min-special requires --special or --special-chars")
	}
	if minSpecial > length {
		return nil, fmt.Errorf("--min-special (%d) cannot exceed password length (%d)", minSpecial, length)
	}
	passwords := make([]string, count)
	for j := 0; j < count; j++ {
		password := make([]byte, length)
		for i := range password {
			password[i] = cs[rand.Intn(len(cs))]
		}
		// guarantee minSpecial special chars at distinct random positions
		positions := rand.Perm(length)[:minSpecial]
		for _, pos := range positions {
			password[pos] = specialChars[rand.Intn(len(specialChars))]
		}
		passwords[j] = string(password)
	}
	return passwords, nil
}
