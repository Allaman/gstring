package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

func splitLines(input string) []string {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
