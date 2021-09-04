/* Package strain implements the `keep` and `discard` operation on collections.
Given a collection and a predicate on the collection's elements, keep
returns a new collection containing those elements where the predicate is
true, while discard returns a new collection containing those elements where
the predicate is false.

This package essentially performs a filter/reject on collections*/
package strain

// Don't understand why Ints{nil} should return nil for both Ints.Keep() and
// Ints.Discard() -- based on test cases
type Ints []int

func (si Ints) Keep(f func(int) bool) Ints {
	var result Ints
	for _, v := range si {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func (si Ints) Discard(f func(int) bool) Ints {
	return si.Keep(func(i int) bool {
		return !f(i)
	})
}

type Strings []string

func (ss Strings) Keep(f func(string) bool) Strings {
	var result Strings
	for _, v := range ss {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

type Lists [][]int

func (ss Lists) Keep(f func([]int) bool) Lists {
	var result Lists
	for _, v := range ss {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
