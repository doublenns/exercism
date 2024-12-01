package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
	switch card {
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
	case "ace":
		value = 11
	default:
		value = 0
	}

	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	decisions := map[string]string{
		"stand": "S",
		"hit":   "H",
		"split": "P",
		"win":   "W",
	}

	myHandValue, dealerCardValue := ParseCard(card1)+ParseCard(card2), ParseCard(dealerCard)
	var decision string

	switch {
	case card1 == "ace" && card2 == "ace":
		decision = decisions["split"]
	case myHandValue == 21:
		if ParseCard(dealerCard) == 11 || ParseCard(dealerCard) == 10 {
			decision = decisions["stand"]
		} else {
			decision = decisions["win"]
		}
	case myHandValue >= 17 && myHandValue <= 20:
		decision = decisions["stand"]
	case myHandValue >= 12 && myHandValue <= 16:
		if dealerCardValue >= 7 {
			decision = decisions["hit"]
		} else {
			decision = decisions["stand"]
		}
	case myHandValue <= 11:
		decision = decisions["hit"]
	}
	return decision
}
