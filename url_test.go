package main

import (
	"testing"
)

func TestURLEncode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "hello+world"},
		{"hello world! @#$%^&*()", "hello+world%21+%40%23%24%25%5E%26%2A%28%29"},
		{"", ""},
		{"abc", "abc"},
		{"Go is awesome!", "Go+is+awesome%21"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := encodeURL(tt.input)
			if result != tt.expected {
				t.Errorf("urlEncode(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestURLDecode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"hello+world", "hello world", false},
		{"hello+world%21+%40%23%24%25%5E%26%2A%28%29", "hello world! @#$%^&*()", false},
		{"", "", false},
		{"abc", "abc", false},
		{"Go+is+awesome%21", "Go is awesome!", false},
		{"%zz", "", true}, // Invalid percent-encoding
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := decodeURL(tt.input)
			if (err != nil) != tt.err {
				t.Fatalf("urlDecode(%q) returned error %v, want error %v", tt.input, err, tt.err)
			}
			if result != tt.expected {
				t.Errorf("urlDecode(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
