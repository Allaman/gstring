package main

import (
	"fmt"
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

func printOutput(s interface{}, trim bool) {
	if trim {
		fmt.Print(s)
	} else {
		fmt.Println(s)
	}
}
