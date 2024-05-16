package main

import (
	"math/rand"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type caseOperation int

const (
	Lower caseOperation = iota
	Upper
)

func getCaseOperation(op caseOperation) func(rune) rune {
	switch op {
	case Lower:
		return func(r rune) rune {
			return unicode.ToLower(r)
		}
	case Upper:
		return func(r rune) rune {
			return unicode.ToUpper(r)
		}
	default:
		return nil
	}
}

func toRandomCase(str string) string {
	var builder strings.Builder
	for _, char := range str {
		r := rand.Intn(2)
		builder.WriteString(string(getCaseOperation(caseOperation(r))(char)))
	}
	return builder.String()
}

func toCamelCase(s string) string {
	words := strings.Fields(s)

	titleCaser := cases.Title(language.Und)

	for i, word := range words {
		if i == 0 {
			words[i] = strings.ToLower(word)
		} else {
			words[i] = titleCaser.String(strings.ToLower(word))
		}
	}

	return strings.Join(words, "")
}

func toSnakeCase(s string) string {
	var result []rune

	for i, r := range s {
		if unicode.IsUpper(r) {
			// Add an underscore before uppercase letters, except at the beginning
			if i > 0 && (unicode.IsLower(rune(s[i-1])) || unicode.IsDigit(rune(s[i-1])) || (i < len(s)-1 && unicode.IsLower(rune(s[i+1])))) {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else if unicode.IsSpace(r) || r == '-' || r == '_' || r == '.' {
			if len(result) > 0 && result[len(result)-1] != '_' {
				result = append(result, '_')
			}
		} else {
			result = append(result, r)
		}
	}

	snake := string(result)

	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	snake = re.ReplaceAllString(snake, "_")

	snake = strings.ToLower(snake)

	snake = strings.Trim(snake, "_")

	return snake
}
