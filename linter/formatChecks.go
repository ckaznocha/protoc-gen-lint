package linter

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func isCamelCase(s string) bool {
	first, size := utf8.DecodeRuneInString(s)

	var numberCount int

	for _, c := range s[size:] {
		if unicode.IsNumber(c) {
			numberCount++
		}
	}

	allNumeric := len(s[size:]) == numberCount

	if unicode.IsLower(first) ||
		(!allNumeric && s == strings.ToUpper(s)) ||
		strings.Contains(s, "_") {
		return false
	}

	return true
}

func isLowerUnderscore(s string) bool {
	return s == strings.ToLower(s)
}

func isUpperUnderscore(s string) bool {
	return s == strings.ToUpper(s)
}
