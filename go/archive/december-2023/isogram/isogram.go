package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	seen := make(map[rune]struct{})

	for _, r := range word {
		if unicode.IsLetter(r) {
			if _, ok := seen[r]; ok {
				return false
			} else {
				seen[r] = struct{}{}
			}
		}
	}
	return true
}
