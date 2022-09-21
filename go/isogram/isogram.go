package isogram

import "unicode"

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, r := range word {
		if unicode.IsLetter(r) {
			if _, ok := seen[unicode.ToLower(r)]; ok {
				return false
			} else {
				seen[unicode.ToLower(r)] = true
			}
		}
	}
	return true
}
