// Package accumulate performs an operation on each element of a given
// collection. It is essentially a []strings.Map() function.
package accumulate

// Accumulate returns a new slice containing the results of applying the
// function f to each string in the original slice.
func Accumulate(ss []string, f func(string) string) []string {
	m := make([]string, len(ss))
	for i, v := range ss {
		m[i] = f(v)
	}
	return m
}
