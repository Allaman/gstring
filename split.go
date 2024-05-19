package main

import (
	"strings"
)

func splitString(s, sep string) []string {
	parts := strings.Split(s, sep)
	nonEmptyParts := []string{}
	for _, part := range parts {
		if part != "" {
			nonEmptyParts = append(nonEmptyParts, part)
		}
	}
	return nonEmptyParts
}

func formatSplittedString(ss []string) string {
	var formated strings.Builder
	for s := range ss {
		formated.Write([]byte(ss[s]))
		if s != len(ss)-1 { // omit last CR
			formated.Write([]byte("\n"))
		}
	}
	return formated.String()
}
