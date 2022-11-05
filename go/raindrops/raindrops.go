package raindrops

import "fmt"

var conversions = []struct {
	factor int
	sound  string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) string {
	var result string

	for _, c := range conversions {
		if number%c.factor == 0 {
			result += c.sound
		}
	}

	if result == "" {
		return fmt.Sprint(number)
	}
	return result
}
