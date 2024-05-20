package main

import (
	"strings"
	"testing"
)

func TestSortLines(t *testing.T) {
	tests := []struct {
		input            string
		desc             bool
		ignoreEmptyLines bool
		unique           bool
		expected         string
	}{
		{
			input: `banana
apple
cherry
banana
apple
date`,
			desc:             false,
			ignoreEmptyLines: true,
			unique:           true,
			expected: `apple
banana
cherry
date`,
		},
		{
			input: `banana
apple
cherry
banana
apple
date`,
			desc:             true,
			ignoreEmptyLines: true,
			unique:           true,
			expected: `date
cherry
banana
apple`,
		},
		{
			input: `a
b

a
b`,
			desc:             false,
			ignoreEmptyLines: true,
			unique:           true,
			expected: `a
b`,
		},
		{
			input: `a
b

a
b`,
			desc:             true,
			ignoreEmptyLines: true,
			unique:           true,
			expected: `b
a`,
		},
		{
			input: `a
b

a
b`,
			desc:             false,
			ignoreEmptyLines: false,
			unique:           false,
			expected: `

a
a
b
b`,
		},
		{
			input:            `one line`,
			desc:             false,
			ignoreEmptyLines: true,
			unique:           true,
			expected:         `one line`,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := sortLines(tt.input, tt.desc, tt.ignoreEmptyLines, tt.unique)
			if strings.TrimSpace(result) != strings.TrimSpace(tt.expected) {
				t.Errorf("sortLines(%q, %t, %t, %t) = %q; want %q", tt.input, tt.desc, tt.ignoreEmptyLines, tt.unique, result, tt.expected)
			}
		})
	}
}
