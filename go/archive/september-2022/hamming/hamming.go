/*
Package hamming provides a handy comparison method to compare 2 strands of
DNS and count the differences between them to see how many mistakes occurred
during cell division.
*/
package hamming

import (
	"errors"
	"strings"
)

// Original implementation used a for loop for iteration
// func Distance(a, b string) (int, error) {
// 	string_length := len(a)
// 	if string_length != len(b) {
// 		return 0, errors.New("the Hamming distance is only defined for sequences of equal length")
// 	}

// 	hamming_distance := 0
// 	for i := 0; i < string_length; i++ {
// 		if a[i] != b[i] {
// 			hamming_distance++
// 		}
// 	}
// 	return hamming_distance, nil
// }

// Implemented later using recursion for practice

// Distance calculates the Hamming Distance between two DNS strands.
// Don't know why `go doc` doesn't show this comment for the function.
func Distance(a, b string) (int, error) {
	var hamming_distance int

	if strings.Compare(a, b) == 0 {
		return 0, nil
	}
	string_length := len(a)
	if string_length != len(b) {
		return 0, errors.New("the Hamming distance is only defined for sequences of equal length")
	}

	// Closure has access to above vars w/o passing them as params
	var compare_runes func(string, string, int) int
	compare_runes = func(a, b string, i int) int {
		if a[i] != b[i] {
			hamming_distance++
		}
		if i < string_length-1 {
			hamming_distance = compare_runes(a, b, i+1)
		}
		return hamming_distance
	}

	return compare_runes(a, b, 0), nil
}
