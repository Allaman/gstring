package main

import (
	"runtime"
	"strings"
)

func head(input string, n int) string {
	lines := splitLines(input)
	if n > len(lines) {
		n = len(lines)
	}
	lineEnding := "\n"
	if runtime.GOOS == "windows" {
		lineEnding = "\r\n"
	}
	return strings.Join(lines[:n], lineEnding)
}
