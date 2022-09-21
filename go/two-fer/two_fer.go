package twofer

import "fmt"

// ShareWith returns a 2-fer sentence of
// "One for <name>, one for me."
// Except if the name is missing, it returns
// "One for you, one for me"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
