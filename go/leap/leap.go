// Package leap is used to determine whether a given year is a leap year or not.
package leap

// IsLeapYear returns a boolean indiciating whether year i is a leap year or not.
func IsLeapYear(y int) bool {
	if y%400 == 0 {
		return true
	} else if y%100 == 0 {
		return false
	} else if y%4 == 0 {
		return true
	}
	return false
}
