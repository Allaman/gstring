package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestGenerateBytes(t *testing.T) {
	tests := []struct {
		count    int
		expected string
	}{
		{0, ""},
		{1, "L"},
		{5, "Lorem"},
		{10, "Lorem ipsu"},
		{864, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ultrices sagittis orci a scelerisque purus. At volutpat diam ut venenatis tellus in. Mattis rhoncus urna neque viverra justo nec ultrices dui sapien. Tellus orci ac auctor augue mauris. Eu scelerisque felis imperdiet proin fermentum. Tortor id aliquet lectus proin nibh nisl condimentum id. Id donec ultrices tincidunt arcu non sodales. Ultrices dui sapien eget mi proin. Bibendum neque egestas congue quisque egestas diam. Sem fringilla ut morbi tincidunt augue interdum. Vel risus commodo viverra maecenas accumsan lacus vel facilisis volutpat. Morbi blandit cursus risus at ultrices mi tempus. Adipiscing vitae proin sagittis nisl rhoncus mattis rhoncus. Sapien pellentesque habitant morbi tristique senectus et netus et malesuada. Lorem"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.count), func(t *testing.T) {
			result := generateBytes(tt.count)
			if result != tt.expected {
				t.Errorf("generateBytes(%d) = %q; want %q", tt.count, result, tt.expected)
			}
		})
	}
}

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
		{"bytes", 5, "Lorem"},
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
