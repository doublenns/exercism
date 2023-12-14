package anagram

import (
	"fmt"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var result []string
	// Make the subject word all lowercase to make character checking easier
	lSub := strings.ToLower(subject)

	check := map[rune]int{}
	for _, r := range lSub {
		check[r]++
	}

	for _, w := range candidates {
		// Also make the candidate lower-cased for comparisons to the subject
		lw := strings.ToLower(w)
		// Initial validation to see if can skip deeper comparisons
		if len(lw) != len(lSub) || lw == lSub {
			continue
		}

		// If the candidate word passes initial validation, turn it into a map
		// & compare the resulting map to that of the subject
		compare := map[rune]int{}
		for _, r := range lw {
			compare[r]++
		}
		if fmt.Sprint(check) == fmt.Sprint(compare) {
			result = append(result, w)
		}
	}

	return result
}
