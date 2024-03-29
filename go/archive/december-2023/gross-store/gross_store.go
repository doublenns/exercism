package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if num, ok := units[unit]; ok {
		bill[item] += num
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// Check if customer supplied correct unit in plain English terms
	if num, ok := units[unit]; ok {
		// Check if the item is even in the bill
		if quantity, ok := bill[item]; ok {
			// Can't remove more items than currently billed for
			if quantity < num {
				return false
			} else {
				bill[item] -= num
				// Simplify end bill by removing all items w/ quanitty of 0
				if bill[item] == 0 {
					delete(bill, item)
				}
				return true
			}
		}
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if quantity, ok := bill[item]; ok {
		return quantity, true
	}
	return 0, false
}
