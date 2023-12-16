package blackjack

var tenCards = map[string]bool{
	"ten":   true,
	"jack":  true,
	"queen": true,
	"king":  true,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var result int

	switch {
	case card == "ace":
		result = 11
	case card == "two":
		result = 2
	case card == "three":
		result = 3
	case card == "four":
		result = 4
	case card == "five":
		result = 5
	case card == "six":
		result = 6
	case card == "seven":
		result = 7
	case card == "eight":
		result = 8
	case card == "nine":
		result = 9
	case tenCards[card]:
		result = 10
	default:
		result = 0
	}
	return result
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var (
		result string

		strategy = map[string]string{
			"Stand": "S",
			"Hit":   "H",
			"Split": "P",
			"Win":   "W",
		}
	)

	switch handTotal := ParseCard(card1) + ParseCard(card2); {
	case handTotal == 22:
		result = strategy["Split"]
	case handTotal == 21:
		// if !tenCards[dealerCard] && dealerCard != "ace" {
		if _, ok := tenCards[dealerCard]; !ok && dealerCard != "ace" {
			result = strategy["Win"]
		} else {
			result = strategy["Stand"]
		}
	case handTotal >= 17 && handTotal <= 20:
		result = strategy["Stand"]
	case handTotal >= 12 && handTotal <= 16:
		if ParseCard(dealerCard) >= 7 {
			result = strategy["Hit"]
		} else {
			result = strategy["Stand"]
		}
	case handTotal <= 11:
		result = strategy["Hit"]
	}
	return result
}
