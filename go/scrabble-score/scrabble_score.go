// Package scrabble generates the scrabble score of a given word.
package scrabble

import (
	"unicode"
)

// letterValue returns the Scrabble score for a given letter.
func letterValue(char rune) (value int) {
	switch unicode.ToUpper(char) {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		value = 1
	case 'D', 'G':
		value = 2
	case 'B', 'C', 'M', 'P':
		value = 3
	case 'F', 'H', 'V', 'W', 'Y':
		value = 4
	case 'K':
		value = 5
	case 'J', 'X':
		value = 8
	case 'Q', 'Z':
		value = 10
	}

	return
}

// Score returns the Scrabble score from a given word.
func Score(word string) int {
	var result int

	for _, r := range word {
		result += letterValue(r)
	}

	return result
}
