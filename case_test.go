package main

import (
	"testing"
	"unicode"
)

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "helloWorld"},
		{"Hello World", "helloWorld"},
		{"hello_world", "helloWorld"},
		{"HELLO WORLD", "helloWorld"},
		{"hello WORLD", "helloWorld"},
		{"", ""},
		{"a", "a"},
		{"this is a test", "thisIsATest"},
		{"convert this to camel case", "convertThisToCamelCase"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := toCamelCase(tt.input)
			if result != tt.expected {
				t.Errorf("toCamelCase(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Test getCaseOperation function
func TestGetCaseOperation(t *testing.T) {
	tests := []struct {
		name     string
		op       caseOperation
		input    rune
		expected rune
	}{
		{"Lowercase conversion", Lower, 'A', 'a'},
		{"Uppercase conversion", Upper, 'a', 'A'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := getCaseOperation(tt.op)
			if fn == nil {
				t.Fatalf("getCaseOperation(%v) returned nil", tt.op)
			}
			result := fn(tt.input)
			if result != tt.expected {
				t.Errorf("getCaseOperation(%v)(%q) = %q; want %q", tt.op, tt.input, result, tt.expected)
			}
		})
	}
}

func TestRandomCase(t *testing.T) {
	original := "Hello, World!"
	randomized := toRandomCase(original)

	if len(randomized) != len(original) {
		t.Fatalf("randomCase(%q) = %q; lengths differ", original, randomized)
	}

	// Convert randomized string to a slice of runes for comparison
	randomizedRunes := []rune(randomized)

	// Check that each character in the original has been either lowercased or uppercased
	for i, char := range original {
		if !unicode.IsLetter(char) {
			// Non-letter characters should remain the same
			if randomizedRunes[i] != char {
				t.Errorf("randomCase(%q) changed non-letter character %q to %q", original, char, randomizedRunes[i])
			}
			continue
		}

		// Letter characters should be either uppercased or lowercased
		lower := unicode.ToLower(char)
		upper := unicode.ToUpper(char)
		if randomizedRunes[i] != lower && randomizedRunes[i] != upper {
			t.Errorf("randomCase(%q) changed %q to %q, which is not a valid case conversion", original, char, randomizedRunes[i])
		}
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World", "hello_world"},
		{"HelloWorld", "hello_world"},
		{"HELLO WORLD", "hello_world"},
		{"hello WORLD", "hello_world"},
		{"", ""},
		{"a", "a"},
		{"thisIsATest", "this_is_a_test"},
		{"convert this to snake case", "convert_this_to_snake_case"},
		{"Test123WithNumbers", "test123_with_numbers"},
		{"Already_Snake_Case", "already_snake_case"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := toSnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("toSnakeCase(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
