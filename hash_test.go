package main

import (
	"testing"
)

func TestCalculateSHA256(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello, World!", "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f"},
		{"Hello, 世界", "a281e84c7f61393db702630c2a6807e871cd3b6896c9e56e22982d125696575c"},
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"abc", "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
		{"Go is awesome!", "d557c06d48fd26fa66dfc2c327288fe815f537addfde447da9e70ae69ceae437"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := calculateSHA256(tt.input)
			if result != tt.expected {
				t.Errorf("calculateSHA256(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCalculateSHA512(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello, World!", "374d794a95cdcfd8b35993185fef9ba368f160d8daf432d08ba9f1ed1e5abe6cc69291e0fa2fe0006a52570ef18c19def4e617c33ce52ef0a6e5fbe318cb0387"},
		{"Hello, 世界", "ab96e79129b670241b07fe92d135dd3f907a38a5d4b36727dc9f471a023efe357ebc884a16e3f536f3184389f20798177f722f8ac4c697eac785fb9908738b1b"},
		{"", "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"},
		{"abc", "ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f"},
		{"Go is awesome!", "32166e6fa484792517898ffef1a33974989420780c703ebd0ad3a67d115d147ddb378f4c765d9bf5788bd1f4507092d0ba746aec900f1f870f76cec6e424bd27"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := calculateSHA512(tt.input)
			if result != tt.expected {
				t.Errorf("calculateSHA512(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCalculateMD5(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "5eb63bbbe01eeed093cb22bb8f5acdc3"},
		{"Hello, 世界", "3dbca55819ed79f62e6f770eef640eee"},
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
		{"abc", "900150983cd24fb0d6963f7d28e17f72"},
		{"Go is awesome!", "572a27a144f77f4a657130b57cdf1742"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := calculateMD5(tt.input)
			if result != tt.expected {
				t.Errorf("calculateMD5(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
