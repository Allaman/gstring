package main

import (
	"testing"
)

func TestEncodeBase64(t *testing.T) {
	original := "Hello, 世界"
	expected := "SGVsbG8sIOS4lueVjA=="

	encoded := encodeBase64(original)
	if encoded != expected {
		t.Errorf("encodeBase64(%q) = %q; want %q", original, encoded, expected)
	}
}

func TestDecodeBase64(t *testing.T) {
	encoded := "SGVsbG8sIOS4lueVjA=="
	expected := "Hello, 世界"

	decoded, err := decodeBase64(encoded)
	if err != nil {
		t.Fatalf("decodeBase64(%q) returned error: %v", encoded, err)
	}
	if decoded != expected {
		t.Errorf("decodeBase64(%q) = %q; want %q", encoded, decoded, expected)
	}
}

func TestDecodeBase64_InvalidInput(t *testing.T) {
	invalid := "SGVsbG8sIOS4lueVjA" // Missing padding

	_, err := decodeBase64(invalid)
	if err == nil {
		t.Errorf("decodeBase64(%q) = expected error, got none", invalid)
	}
}
