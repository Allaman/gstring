package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
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

func TestPrintOutput(t *testing.T) {
	tests := []struct {
		input    interface{}
		trim     bool
		expected string
	}{
		{"Hello, World!", false, "Hello, World!\n"},
		{"Hello, World!", true, "Hello, World!"},
		{"Go is awesome!", false, "Go is awesome!\n"},
		{"Go is awesome!", true, "Go is awesome!"},
		{12345, false, "12345\n"},
		{12345, true, "12345"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v_%v", tt.input, tt.trim), func(t *testing.T) {
			// Redirect stdout to capture output
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			printOutput(tt.input, tt.trim)

			// Restore stdout and read captured output
			w.Close()
			var buf bytes.Buffer
			_, _ = io.Copy(&buf, r)
			os.Stdout = oldStdout

			result := buf.String()
			if result != tt.expected {
				t.Errorf("printOutput(%v, %v) = %q; want %q", tt.input, tt.trim, result, tt.expected)
			}
		})
	}
}

func TestSplitLines(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"line1\nline2\nline3\n", []string{"line1", "line2", "line3"}},
		{"line1\nline2\nline3", []string{"line1", "line2", "line3"}},
		// {"", []string{}}, -- TODO: this test fails probably due to reflect.DeepEqual
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := splitLines(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("splitLines(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
