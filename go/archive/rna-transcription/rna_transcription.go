// Package strand returns the RNA complement of a given DNA strand
package strand

import (
	"strings"
)

// ToRNA replaces each DNA (string) nucleotide with it's complement.
func ToRNA(dna string) string {
	r := strings.NewReplacer(
		"G", "C",
		"C", "G",
		"T", "A",
		"A", "U",
	)
	return r.Replace(dna)
}
