// Package strutil provides string utility functions as an example of a public pkg/ package.
package strutil

import "strings"

// Reverse reverses a string by Unicode code points.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsBlank reports whether a string is empty or contains only whitespace.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}
