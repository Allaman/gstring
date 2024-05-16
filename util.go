package main

import (
	"io"
	"os"
)

func readFromSTDIN() (string, error) {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}

	return string(stdin), nil
}
