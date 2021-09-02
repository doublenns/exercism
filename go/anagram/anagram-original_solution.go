// Package Anagram detects whether words are anagrams of each other.
// An anagram is a rearrangement of letters to form a new word.
// BenchmarkDetectAnagrams-8          64774             19440 ns/op            8786 B/op        132 allocs/op
package anagram

import (
	"strings"
)

// stringToMap takes in a string and returns a map of the specified string,
// where the keys are each unique rune contained in the string and the values
// are the number of times that rune occurs in the string.
func stringToMap(s string) map[rune]int {
	result := make(map[rune]int)
	for _, v := range s {
		_, found := result[v]
		if found {
			result[v]++
		} else {
			result[v] = 0
		}
	}
	return result
}

type Strings []string

// keep implements the `keep` operation on a `Strings` ( []string ) collection.
// Given a collection and a predicate on the collection's elements, keep
// returns a new collection containing those elements where the predicate is
// true.
func (ss Strings) keep(f func(string) bool) Strings {
	result := make([]string, 0)
	for _, v := range ss {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Detect returns the sublist of a list of candidate words that are anagrams
// of a given word s.
func DetectOriginal(s string, candidates []string) []string {
	wordMap := stringToMap(strings.ToLower(s))

	// Using closure so that func automatically has access too all vars passed into
	// the detect func, plus the `wordMap` var above, without needing to explicitly
	// pass the values in.
	compare := func(c string) bool {
		if strings.EqualFold(c, s) { // words can't be anagrams of themselves
			return false
		}
		candidateMap := stringToMap(strings.ToLower(c))
		if len(wordMap) != len(candidateMap) { // Must include the same # of unique runes
			return false
		}
		for k, v := range wordMap {
			if w, ok := candidateMap[k]; !ok || v != w { // Must have same # of each rune
				return false
			}
		}
		return true
	}

	c := append(make(Strings, 0, len(candidates)), candidates...)
	// `keep()` returns a new collection instead of modifying input in-place
	result := c.keep(compare)

	return result
}
