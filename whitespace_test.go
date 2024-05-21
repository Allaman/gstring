package main

import (
	"testing"
)

func TestRemoveWhitespace(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		removeSpaces     bool
		removeTabs       bool
		removeCR         bool
		removeNewlines   bool
		removeEmptyLines bool
		expected         string
	}{
		{
			name:             "Remove spaces",
			input:            "Hello, World!",
			removeSpaces:     true,
			removeTabs:       false,
			removeCR:         false,
			removeNewlines:   false,
			removeEmptyLines: false,
			expected:         "Hello,World!",
		},
		{
			name:             "Remove tabs",
			input:            "Hello,\tWorld!",
			removeSpaces:     false,
			removeTabs:       true,
			removeCR:         false,
			removeNewlines:   false,
			removeEmptyLines: false,
			expected:         "Hello,World!",
		},
		{
			name:             "Remove carriage returns",
			input:            "Hello,\rWorld!",
			removeSpaces:     false,
			removeTabs:       false,
			removeCR:         true,
			removeNewlines:   false,
			removeEmptyLines: false,
			expected:         "Hello,World!",
		},
		{
			name:             "Remove newlines",
			input:            "Hello,\nWorld!",
			removeSpaces:     false,
			removeTabs:       false,
			removeCR:         false,
			removeNewlines:   true,
			removeEmptyLines: false,
			expected:         "Hello,World!",
		},
		{
			name:             "Remove all whitespace characters",
			input:            "Hello,\tWorld!\r\nThis is a test.\nWith multiple lines\tand spaces.",
			removeSpaces:     true,
			removeTabs:       true,
			removeCR:         true,
			removeNewlines:   true,
			removeEmptyLines: false,
			expected:         "Hello,World!Thisisatest.Withmultiplelinesandspaces.",
		},
		{
			name:             "Remove empty lines",
			input:            "Hello,\tWorld!\r\nThis is a test.\n\nWith multiple lines\tand spaces.\n\n   \nAnother line.",
			removeSpaces:     false,
			removeTabs:       false,
			removeCR:         false,
			removeNewlines:   false,
			removeEmptyLines: true,
			expected:         "Hello,\tWorld!\r\nThis is a test.\nWith multiple lines\tand spaces.\nAnother line.",
		},
		{
			name:             "Remove all whitespace characters and empty lines",
			input:            "Hello,\tWorld!\r\nThis is a test.\n\nWith multiple lines\tand spaces.\n\n   \nAnother line.",
			removeSpaces:     true,
			removeTabs:       true,
			removeCR:         true,
			removeNewlines:   true,
			removeEmptyLines: true,
			expected:         "Hello,World!Thisisatest.Withmultiplelinesandspaces.Anotherline.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeWhitespace(tt.input, tt.removeSpaces, tt.removeTabs, tt.removeCR, tt.removeNewlines, tt.removeEmptyLines)
			if result != tt.expected {
				t.Errorf("Expected '%s', but got '%s'", tt.expected, result)
			}
		})
	}
}
