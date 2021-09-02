package twofer

import "fmt"

// Takes in a name and outputs "One for <name>, one for me." but uses "you" if the name is absent
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
