// Package bob simulates how Bob, a lackadaisical teenager with very limited
// responses, would respond to remarks made in a conversation.
package bob

import (
	"strings"
	"unicode"
)

// isUpper checks if string only contains upper-cased letters.
// The string needs to actually contain letters to be considered upper-cased.
func isUpper(s string) bool {
	letterCount := 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			letterCount++
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return letterCount > 0
}

// isQuestion checks if a string ends with a question mark
func isQuestion(s string) bool {
	// return s[len(s)-1] == '?'
	// Using standard library instead of manually checking element at last
	// index of string:
	// func HasSuffix(s, suffix string) bool {
	//     return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
	// }
	return strings.HasSuffix(s, "?")
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Remove all leading & trailing whitespace from string before processing
	remark = strings.TrimSpace(remark)

	var response string
	switch {
	// After refactoring implementation of isQuestion(), no longer need to
	// check for empty remark first to prevent out of range index panic from
	// using isQuestion() on an empty string -- this case can be placed
	// anywhere in the switch now
	case remark == "":
		response = "Fine. Be that way!"
	case isUpper(remark) && isQuestion(remark):
		response = "Calm down, I know what I'm doing!"
	case isUpper(remark):
		response = "Whoa, chill out!"
	case isQuestion(remark):
		response = "Sure."
	default:
		response = "Whatever."
	}
	return response
}
