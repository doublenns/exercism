// Package proverb generates a relevant proverb, given a list of input strings.
package proverb

import "fmt"

// Proverb processes a list of input strings, 2 strings at a time, and returns a relevant proverb.
func Proverb(rhyme []string) []string {
	var result []string
	if len(rhyme) > 0 {
		initialWant := rhyme[0]
		for {
			if len(rhyme) < 2 {
				result = append(result, fmt.Sprintf("And all for the want of a %s.", initialWant))
				break
			}
			var s []string
			s, rhyme = rhyme[0:2], rhyme[1:] // Process 2 items of array/queue but only remove 1 item
			result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", s[0], s[1]))
		}
	}
	return result
}
