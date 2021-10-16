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
	if (ParseCard(card1) + ParseCard(card2)) == 21 {
		return true
	}
	return false
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	switch {
	case isBlackjack == false:
		return "P"
	case isBlackjack == true && dealerScore < 10:
		return "W"
	// Cleaner to show the actual boolean statement than using a default statement
	// However, compiler complains about not having a default return
	// case isBlackjack == true && dealerScore >= 10:
	default:
		return "S"
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	switch {
	case handScore >= 17:
		return "S"
	case handScore <= 11:
		return "H"
	case dealerScore >= 7:
		return "H"
	// case dealerScore <7:
	default:
		return "S"
	}
}
