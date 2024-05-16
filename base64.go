package main

import "encoding/base64"

func encodeBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func decodeBase64(str string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
