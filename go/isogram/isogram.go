// Package isogram checks whether a string is an isogram, which is a word or
// phrase without a repeating letter (although spaces and hyphens are allowed
// to appear multiple times.)
package isogram

import (
	"strings"
)

// removeNonAlphas returns s after removing all non-alphanumeric characters,
// including hyphens and spaces. The input string is already lowercased, so
// this func only checks for lowercased characters.
func removeNonAlphas(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if 'a' <= b && b <= 'z' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

// IsIsogram returns a boolean based on whether input string s is an isogram
// or not.
func IsIsogram(s string) bool {
	s = strings.ToLower(s)
	s = removeNonAlphas(s)

	seenFlag := map[rune]bool{}
	for _, c := range s {
		if seenFlag[c] {
			return false
		}
		seenFlag[c] = true
	}
	return true
}
