/*
BenchmarkDetectAnagrams-8          45704             25236 ns/op            6848 B/op        155 allocs/op

Importing `sort` and using `sort.Strings()` is an expensive operation and is
less "performant" than my original solution, which compares 2 instances of
`map[rune]int`. This solution, however, is much, much easier to read -- but
I'm attributing much of that to the "naked" append in the `Detect()` function
in lieu of using the "filter/strain" I created in the `Keep()` withhin my OG
solution. (I mostly implemented Keep() for learning purposes in the OG)
*/
package anagram

import (
	"sort"
	"strings"
)

func sorted(in string) string {
	chars := strings.Split(in, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func Detect(s string, candidates []string) []string {
	anagrams := []string{}
	compare := sorted(strings.ToLower(s))

	for _, candidate := range candidates {
		check := sorted(strings.ToLower(candidate))
		if !strings.EqualFold(s, candidate) && check == compare {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}
