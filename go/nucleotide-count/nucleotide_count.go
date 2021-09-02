// Package dna compute how many times each nucleotide occurs in a single stranded DNA string.
package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides, represented as a string.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	h['A'] = strings.Count(string(d), "A")
	h['C'] = strings.Count(string(d), "C")
	h['G'] = strings.Count(string(d), "G")
	h['T'] = strings.Count(string(d), "T")

	if h['A']+h['C']+h['G']+h['T'] < len(d) {
		return h, errors.New("the dna sequence contains invalid nucleotides")
	}
	return h, nil
}
