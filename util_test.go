package main

import (
	"os"
	"testing"
)

func TestReadFromSTDIN(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		{"Hello, World!", "Hello, World!"},
		{"Hello, 世界", "Hello, 世界"},
		{"", ""},
		{"This is a test.", "This is a test."},
	}

	for _, tt := range tests {
		originalStdin := os.Stdin

		// Restore original os.Stdin after the test
		defer func() { os.Stdin = originalStdin }()

		// Set os.Stdin to a bytes.Buffer with test input
		r, w, _ := os.Pipe()
		_, _ = w.Write([]byte(tt.input))
		w.Close()
		os.Stdin = r

		t.Run(tt.input, func(t *testing.T) {
			result, err := readFromSTDIN()
			if err != nil {
				t.Fatalf("Error reading from stdin: %v", err)
			}
			if result != tt.expected {
				t.Errorf("readFromSTDIN() = %q; want %q", result, tt.expected)
			}
		})
	}
}
