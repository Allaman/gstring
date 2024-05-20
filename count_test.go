package main

import (
	"testing"
)

func TestCountBytes(t *testing.T) {
	tests := []struct {
		input    string
		format   bool
		expected string
	}{
		{"Hello, World!", false, "13"},
		{"Hello, 世界", false, "13"},
		{"", false, "0"},
		{"abc", false, "3"},
		{"こんにちは", false, "15"}, // 5 characters, 3 bytes each in UTF-8
		{"Hello, World!", true, "13 B"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := countBytes(tt.input, tt.format)
			if result != tt.expected {
				t.Errorf("countBytes(%q) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCountChars(t *testing.T) {
	tests := []struct {
		input    string
		char     rune
		expected int
	}{
		{"Hello, World!", 0, 13},
		{"Hello, World!", 'o', 2},
		{"Hello, 世界", 0, 9},
		{"Hello, 世界", '世', 1},
		{"", 0, 0},
		{"abc", 0, 3},
		{"abc", 'a', 1},
		{"こんにちは", 0, 5},
		{"こんにちは", 'に', 1},
		{"Go is awesome!", 0, 14},
		{"Go is awesome!", 'o', 2},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			var result int
			if tt.char == 0 {
				result = countChars(tt.input)
			} else {
				result = countChars(tt.input, tt.char)
			}
			if result != tt.expected {
				t.Errorf("countChars(%q, %q) = %d; want %d", tt.input, tt.char, result, tt.expected)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		word     string
		expected int
	}{
		{"hello world this is a test", "", 6},
		{"hello world this is a test", "hello", 1},
		{"hello hello hello", "hello", 3},
		{"", "", 0},
		{"one", "", 1},
		{"one two two three three three", "", 6},
		{"one two two three three three", "three", 3},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if tt.word == "" {
				result := countWords(tt.input)
				if result != tt.expected {
					t.Errorf("countWords(%q) = %d; want %d", tt.input, result, tt.expected)
				}
			} else {
				result := countWords(tt.input, tt.word)
				if result != tt.expected {
					t.Errorf("countWords(%q, %q) = %d; want %d", tt.input, tt.word, result, tt.expected)
				}
			}
		})
	}
}
