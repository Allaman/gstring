package main

import (
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/dustin/go-humanize"
)

func countBytes(s string, human bool) string {
	if human {
		return humanize.Bytes(uint64(len(s)))
	}
	return strconv.Itoa(len(s))
}

func countChars(str string, chars ...rune) int {
	if len(chars) > 0 {
		char := chars[0]
		count := 0
		for _, c := range str {
			if c == char {
				count++
			}
		}
		return count
	}
	return utf8.RuneCountInString(str)
}

func countWords(str string, words ...string) int {
	wordList := strings.Fields(str)

	if len(words) > 0 {
		word := words[0]
		count := 0
		for _, w := range wordList {
			if w == word {
				count++
			}
		}
		return count
	}
	return len(wordList)
}
