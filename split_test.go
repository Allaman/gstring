package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplitString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		sep      string
		expected []string
	}{
		{"Comma separator", "one,two,three", ",", []string{"one", "two", "three"}},
		{"Space separator", "one two three", " ", []string{"one", "two", "three"}},
		{"Newline separator", "one\ntwo\nthree", "\n", []string{"one", "two", "three"}},
		{"Empty separator", "onetwothree", "", []string{"o", "n", "e", "t", "w", "o", "t", "h", "r", "e", "e"}},
		{"No separator match", "onetwothree", ",", []string{"onetwothree"}},
		{"Leading and trailing separators", ",one,two,three,", ",", []string{"one", "two", "three"}},
		{"Multiple spaces", "one  two  three", "  ", []string{"one", "two", "three"}},
		{"Empty string", "", ",", []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitString(tt.input, tt.sep)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("splitString(%q, %q) = %v; want %v", tt.input, tt.sep, result, tt.expected)
			}
		})
	}
}

func TestFormatSplittedString(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"line1", "line2", "line3"}, "line1\nline2\nline3"},
		{[]string{"one", "two", "three", "four"}, "one\ntwo\nthree\nfour"},
		{[]string{"singleline"}, "singleline"},
		{[]string{}, ""},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.input, ","), func(t *testing.T) {
			result := formatSplittedString(tt.input)
			if result != tt.expected {
				t.Errorf("formatSplittedString(%v) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
