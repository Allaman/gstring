package main

import (
	"sort"
	"strings"
)

func sortLines(input string, desc, ignoreEmptyLines, unique bool) string {
	lines := strings.Split(input, "\n")
	if ignoreEmptyLines {
		nonEmptyLines := make([]string, 0, len(lines))
		for _, line := range lines {
			if strings.TrimSpace(line) != "" {
				nonEmptyLines = append(nonEmptyLines, line)
			}
		}
		lines = nonEmptyLines
	}

	if unique {
		lines = uniqueLines(lines)
	}

	if desc {
		sort.Sort(sort.Reverse(sort.StringSlice(lines)))
	} else {
		sort.Strings(lines)
	}
	return strings.Join(lines, "\n")
}

func uniqueLines(lines []string) []string {
	lineMap := make(map[string]struct{})
	uniqueLines := make([]string, 0, len(lines))
	for _, line := range lines {
		if _, exists := lineMap[line]; !exists {
			lineMap[line] = struct{}{}
			uniqueLines = append(uniqueLines, line)
		}
	}
	return uniqueLines
}
