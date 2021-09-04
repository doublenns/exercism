package pangram

import (
	"strings"
)

func removeNonLowerAlphas(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if 'a' <= b && b <= 'z' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func IsPangram(s string) bool {
	letters := make(map[rune]bool)

	for _, r := range removeNonLowerAlphas(strings.ToLower(s)) {
		letters[r] = true
	}
	return len(letters) == 26
}
