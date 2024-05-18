package main

import (
	"encoding/hex"
	"strings"
)

func stringToHex(s string, formated bool) string {
	hexStr := hex.EncodeToString([]byte(s))
	if !formated {
		return hexStr
	}
	var spacedHex strings.Builder
	for i := 0; i < len(hexStr); i += 2 {
		if i > 0 {
			spacedHex.WriteByte(' ')
		}
		spacedHex.WriteString(hexStr[i : i+2])
	}
	return spacedHex.String()
}

func hexToString(hexStr string) (string, error) {
	cleanedHexStr := strings.ReplaceAll(hexStr, " ", "")
	bytes, err := hex.DecodeString(cleanedHexStr)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
