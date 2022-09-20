package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten", "jack", "queen", "king":
		value = 10
	default:
		value = 0
	}

	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	h := ParseCard(card1) + ParseCard(card2)
	var r string

	// Go switch statements don't cascade -- need to explicitly call for a fallthrough
	switch {
	case card1 == "ace" && card2 == "ace":
		r = "P"
	case h == 21:
		if ParseCard(dealerCard) < 10 {
			r = "W"
		} else {
			r = "S"
		}
	case h >= 17:
		r = "S"
	case h >= 12:
		if ParseCard(dealerCard) >= 7 {
			r = "H"
		} else {
			r = "S"
		}
	default:
		r = "H"
	}

	return r
}
