package main

import (
	"runtime"
	"strings"
	"testing"
)

func TestHead(t *testing.T) {
	tests := []struct {
		input    string
		n        int
		expected string
	}{
		{"line1\nline2\nline3\nline4\nline5\n", 2, "line1\nline2"},
		{"line1\nline2\nline3\nline4\nline5\n", 0, ""},
		{"line1\nline2\nline3\nline4\nline5\n", 5, "line1\nline2\nline3\nline4\nline5"},
		{"line1\nline2\nline3\nline4\nline5\n", 10, "line1\nline2\nline3\nline4\nline5"},
		{"line1\nline2\nline3\nline4\nline5", 3, "line1\nline2\nline3"},
		{"", 3, ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			// Determine the expected line ending based on the operating system
			lineEnding := "\n"
			if runtime.GOOS == "windows" {
				lineEnding = "\r\n"
			}

			// Replace expected line endings with the appropriate ones for the OS
			expected := strings.ReplaceAll(tt.expected, "\n", lineEnding)

			result := head(tt.input, tt.n)
			if result != expected {
				t.Errorf("head(%d) = %q; want %q", tt.n, result, expected)
			}
		})
	}
}
