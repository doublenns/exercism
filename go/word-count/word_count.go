package wordcount

import "strings"

type Frequency map[string]int

func replaceSeparators(s string) string {
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

func WordCount(phrase string) Frequency {
	result := Frequency{}
	phrase = replaceSeparators(strings.ToLower(phrase))

	words := strings.Fields(phrase)
	for _, w := range words {
		w = strings.Trim(w, "'")
		if len(w) > 0 {
			result[w]++
		}
	}

	return result
}
