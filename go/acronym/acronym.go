// Package acronym takes a string (phrase) and converts it into an acronym.
package acronym

import (
	"strings"
	"unicode"
)

// RemoveNonAlphas returns s after removing all non-alphanumeric characters,
// retaining hyphens and spaces.
func RemoveNonAlphas(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			b == ' ' ||
			b == '-' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

// SplitWords slices s into all substrings separated by spaces and hyphens and
// returns a slice of the substrings (words) between those separators.
func SplitWords(s string) []string {
	splitDelimiters := func(c rune) bool {
		return c == 45 || unicode.IsSpace(c) // 45 is rune "-"
	}
	return strings.FieldsFunc(s, splitDelimiters)
}

// Abbreviate takes in s (phrase), removes non-alphanumerics, splits
// the string on white spaces and hyphens, then returns the first characters
// of each substring (word) in a single capitalized string.
func Abbreviate(s string) string {

	s = RemoveNonAlphas(s)
	words := SplitWords(s)

	var sb strings.Builder
	for _, v := range words {
		sb.WriteString(string(v[0]))
	}
	return strings.ToUpper(sb.String())
}
