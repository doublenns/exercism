package raindrops

import (
	"fmt"
	"strings"
)

var conversions = []struct {
	factor int
	sound  string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) string {
	var result strings.Builder

	for _, c := range conversions {
		if number%c.factor == 0 {
			result.WriteString(c.sound)
		}
	}

	if result.String() == "" {
		return fmt.Sprint(number)
	}
	return result.String()
}
