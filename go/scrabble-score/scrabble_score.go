package scrabble

import (
	"strings"
)

func letterValue(char rune) (value int) {
	switch char {
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

func Score(word string) int {
	u := strings.ToUpper(word)
	var result int

	for _, r := range u {
		result += letterValue(r)
	}

	return result
}
