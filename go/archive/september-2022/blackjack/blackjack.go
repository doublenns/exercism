package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// Exercise wants to determine this using conditionals instead of a map
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	var choice string
	switch {
	case !isBlackjack:
		choice = "P"
	case dealerScore < 10:
		choice = "W"
	case dealerScore >= 10:
		choice = "S"
	}
	return choice
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	var choice string
	switch {
	case handScore >= 17:
		choice = "S"
	case handScore <= 11:
		choice = "H"
	case dealerScore >= 7:
		choice = "H"
	case dealerScore < 7:
		choice = "S"
	}
	return choice
}
