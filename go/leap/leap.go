// Package leap reports whether a given year is a leap year.
package leap

// IsLeapYear returns whether a given year is a leap year.
func IsLeapYear(y int) bool {
	if y%100 == 0 {
		return y%400 == 0
	}

	return y%4 == 0
}
