// Package twofer helps you determine what to say as I give away a cookie.
package twofer

import "fmt"

// ShareWith returns a string instructing me what to say to my friend as I
// give them a cookie.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
