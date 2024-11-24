package blackjack

var tenCards = map[string]bool{
	"ten":   true,
	"jack":  true,
	"queen": true,
	"king":  true,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
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
	// case tenCards[card]:
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var (
		// Using this `result` variable instead of setting a default switch case
		// in favor of more explicit case statements
		result   string
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
