package blackjack

var tenCards = map[string]bool{
	"ten":   true,
	"jack":  true,
	"queen": true,
	"king":  true,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine":
		return 9
	case tenCards[card]:
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
