package linter

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func isCamelCase(s string) bool {
	first, _ := utf8.DecodeRuneInString(s)
	if unicode.IsLower(first) ||
		s == strings.ToUpper(s) ||
		strings.Contains(s, "_") {
		return false
	}
	return true
}

func isLowerUnderscore(s string) bool {
	if s == strings.ToLower(s) {
		return true
	}
	return false
}

func isUpperUnderscore(s string) bool {
	if s == strings.ToUpper(s) {
		return true
	}
	return false
}
