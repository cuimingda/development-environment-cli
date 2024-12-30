package utils

import "strings"

func TrimmedStringFromBytes(b []byte) string {
	return strings.TrimSpace(string(b))
}
