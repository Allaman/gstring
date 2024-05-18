package main

import "net/url"

func encodeURL(s string) string {
	return url.QueryEscape(s)
}

func decodeURL(s string) (string, error) {
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		return "", err
	}
	return decoded, nil
}
