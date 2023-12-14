// Package raindrops converts a number into a string that contains raindrop
// sounds corresponding to the factors of the provided number.
package raindrops

import "strconv"

func Convert(raindrops int) string {
	factorSounds := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	result := ""
	if raindrops%3 == 0 {
		result += factorSounds[3]
	}
	if raindrops%5 == 0 {
		result += factorSounds[5]
	}
	if raindrops%7 == 0 {
		result += factorSounds[7]
	}
	if result == "" {
		result += strconv.Itoa(raindrops)
	}
	return result
}
