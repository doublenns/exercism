/*
Package wordcount returns the number of occurrences of each word in a given phrase.

For the purposes of this package you can expect that a word will always be one of:
1) A number composed of one or more ASCII digits (ie "0" or "1234") OR
2) simple word composed of one or more ASCII letters (ie "a" or "they") OR
3)  contraction of two simple words joined by a single apostrophe (ie "it's" or "they're")

When counting words this package assumes the following rules:
1) The count is case insensitive (ie "You", "you", and "YOU" are 3 uses of the same word)
2) The count is unordered; the tests will ignore how words and counts are ordered
3) Other than the apostrophe in a contraction all forms of punctuation are ignored
4) The words can be separated by any form of whitespace (ie "\t", "\n", " ")
*/
package wordcount

import "strings"

type Frequency map[string]int

// ReplaceSeparators replaces all characters in a string except for lower-cased
// alphas, digits, and apostraphes.
func ReplaceSeparators(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if 'a' <= b && b <= 'z' ||
			'0' <= b && b <= '9' ||
			b == '\'' {
			result.WriteByte(b)
		} else {
			result.WriteByte(' ')
		}
	}
	return result.String()
}

// WordCount returns the occurrences of every word in a given string.
func WordCount(s string) Frequency {
	result := Frequency{}
	words := strings.Fields(ReplaceSeparators(strings.ToLower(s)))
	for _, word := range words {
		word = strings.Trim(word, "'")
		result[word]++
	}
	return result
}
