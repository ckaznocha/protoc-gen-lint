package linter

import (
	"regexp"
	"strings"
)

// @see https://google.github.io/styleguide/javaguide.html#s5.3-camel-case
var camelCaseRegexPattern, _ = regexp.Compile("\\A([A-Z][a-z0-9]+)((\\d)|([A-Z0-9][a-z0-9]+))*\\z")

func isCamelCase(s string) bool {
	return camelCaseRegexPattern.Match([]byte(s))
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
