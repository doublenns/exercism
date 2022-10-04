package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	check := map[rune]struct{}{}
	for _, r := range strings.ToLower(input) {
		if r >= 'a' && r <= 'z' {
			check[r] = struct{}{}
		}
	}

	return len(check) == 26
}
