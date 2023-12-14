/*
Package partyrobot implements v2.0 of a robot to greet people and help them to their seats.

Once there was an eccentric programmer living in a strange house with barred windows.
One day he accepted a job from an online job board to build a party robot. The robot
is supposed to greet people and help them to their seats. The first edition was very
technical and showed the programmer's lack of human interaction.
Some of which also made it into the next edition.
*/
package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s", name)
}

// HappyBirthday wishes happy birthday to the birthday person and stands out his age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbour string, direction string, distance float64) string {
	return fmt.Sprintf(Welcome(name) +
		"You have been assigned to table %X. Your table is on the %s, exactly %.1f meters from here.\n" +
		"You will be sitting next to %s.",
		name, table, direction, distance, neighbour
)
}
