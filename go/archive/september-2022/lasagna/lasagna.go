// Package lasagna contains functions to help cook a brilliant dish of
// Gopher's Gorgeous Lasagna from your favorite cooking book.
package lasagna

// Overtime returns the expected oven time in minutes.
func OvenTime() int { return 40 }

// RemainingOvenTime calculates the remaining oven time in minutes.
func RemainingOvenTime(t int) int { return OvenTime() - t }

// PreparationTime calculates the preparation time in minutes.
func PreparationTime(l int) int { return l * 2 }

// ElapsedTime calculates the elapsed working/prep time in minutes.
func ElapsedTime(l, t int) int {
	return PreparationTime(l) + t
}
