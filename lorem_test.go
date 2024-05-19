package main

import (
	"strings"
	"testing"
)

func TestGenerateWords(t *testing.T) {
	tests := []struct {
		count    int
		expected string
	}{
		{5, "Lorem ipsum dolor sit amet,"},
		{10, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do"},
		{50, strings.Join(loremIpsumWords[:50], " ")},
	}

	for _, tt := range tests {
		t.Run(strings.Join(loremIpsumWords[:tt.count], " "), func(t *testing.T) {
			result := generateWords(tt.count)
			if result != tt.expected {
				t.Errorf("generateWords(%d) = %q; want %q", tt.count, result, tt.expected)
			}
		})
	}
}

func TestGenerateSentences(t *testing.T) {
	tests := []struct {
		count    int
		expected string
	}{
		{1, loremIpsumSentences[0] + "."},
		{2, strings.Join(loremIpsumSentences[:2], ".") + "."},
		{5, strings.Join(loremIpsumSentences[:5], ".") + "."},
	}

	for _, tt := range tests {
		t.Run(strings.Join(loremIpsumSentences[:tt.count], ".")+".", func(t *testing.T) {
			result := generateSentences(tt.count)
			if result != tt.expected {
				t.Errorf("generateSentences(%d) = %q; want %q", tt.count, result, tt.expected)
			}
		})
	}
}

func TestGenerateParagraphs(t *testing.T) {
	tests := []struct {
		count    int
		expected string
	}{
		{1, loremIpsum},
		{2, loremIpsum + "\n\n" + loremIpsum},
	}

	for _, tt := range tests {
		t.Run(strings.Repeat(loremIpsum+"\n\n", tt.count)[:len(strings.Repeat(loremIpsum+"\n\n", tt.count))-2], func(t *testing.T) {
			result := generateParagraphs(tt.count)
			if result != tt.expected {
				t.Errorf("generateParagraphs(%d) = %q; want %q", tt.count, result, tt.expected)
			}
		})
	}
}

func TestGenerateLoremIpsum(t *testing.T) {
	tests := []struct {
		lengthType string
		count      int
		expected   string
	}{
		{"words", 5, "Lorem ipsum dolor sit amet,"},
		{"sentences", 2, strings.Join(loremIpsumSentences[:2], ".") + "."},
		{"paragraphs", 1, loremIpsum},
	}

	for _, tt := range tests {
		t.Run(tt.lengthType, func(t *testing.T) {
			result := generateLoremIpsum(tt.lengthType, tt.count)
			if result != tt.expected {
				t.Errorf("generateLoremIpsum(%q, %d) = %q; want %q", tt.lengthType, tt.count, result, tt.expected)
			}
		})
	}
}
