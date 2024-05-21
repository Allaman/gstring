package main

import (
	"testing"
)

func TestParsePermissions(t *testing.T) {
	tests := []struct {
		octal    string
		expected string
	}{
		{"000", "---------"},
		{"755", "rwxr-xr-x"},
		{"400", "r--------"},
		{"0000", "---------"},
		{"0755", "rwxr-xr-x"},
		{"0644", "rw-r--r--"},
		{"0700", "rwx------"},
		{"0777", "rwxrwxrwx"},
		{"7400", "r-S--S--T"},
		{"4755", "rwsr-xr-x"},
		{"2745", "rwxr-Sr-x"},
		{"1777", "rwxrwxrwt"},
		{"6777", "rwsrwsrwx"},
		{"4000", "--S------"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result, _ := parseUnixPermissions(tt.octal)
			if result != tt.expected {
				t.Errorf("ParsePermissions(%s) = %s; expected %s", tt.octal, result, tt.expected)
			}
		})
	}
}
