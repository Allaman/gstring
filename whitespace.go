package main

import (
	"regexp"
	"strings"
)

func removeWhitespace(input string, removeSpaces bool, removeTabs bool, removeCR bool, removeNewlines bool, removeEmptyLines bool) string {
	pattern := ""
	if removeSpaces {
		pattern += " "
	}
	if removeTabs {
		pattern += "\t"
	}
	if removeCR {
		pattern += "\r"
	}
	if removeNewlines {
		pattern += "\n"
	}

	// Replace specified whitespace characters
	if pattern != "" {
		re := regexp.MustCompile("[" + regexp.QuoteMeta(pattern) + "]")
		input = re.ReplaceAllString(input, "")
	}

	// Remove empty lines if specified
	if removeEmptyLines {
		lines := strings.Split(input, "\n")
		var nonEmptyLines []string
		for _, line := range lines {
			if strings.TrimSpace(line) != "" {
				nonEmptyLines = append(nonEmptyLines, line)
			}
		}
		input = strings.Join(nonEmptyLines, "\n")
	}

	return input
}
