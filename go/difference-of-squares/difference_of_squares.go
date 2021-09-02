// Package diffsquares finds the difference between the square of the sum and
// the sum of the squares of the first N natural numbers.
package diffsquares

// SquareofSum returns the square of the sum of the first n natural numbers.
func SquareOfSum(n int) int {
	// Faster to use multiplication than math.Pow()
	// This also avoids having to convert to float64 then back to int
	// return (n * n) * (n + 1) * (n + 1) / 4
	// Below formula doesn't use exponents at all, is easier to read, and just as fast
	sum := n * (n + 1) / 2 // http://www.trans4mind.com/personal_development/mathematics/series/sumNaturalNumbers.htm#mozTocId914933
	return sum * sum
}

// SumofSquares returns the sum of the squares of the first ten natural numbers.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6 // http://www.trans4mind.com/personal_development/mathematics/series/sumNaturalSquares.htm#summation
}

// Difference returns the square of the sum of the first n natural numbers and
// the sum of the squares of the first n natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
