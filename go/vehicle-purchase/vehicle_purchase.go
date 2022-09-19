package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	// More lines of code, but more readable.
	var choice string
	if option1 < option2 {
		choice = option1
	} else {
		choice = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", choice)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var discount float64

	// Using a switch statement here is more lines of code than the initial implementation,
	// but ultimately more readable to understand what the code is doing (and maintainable
	// for future discount modifications)).
	switch {
	case age < 3:
		discount = .8
	case age > 3 && age < 10:
		discount = .7
	case age >= 10:
		discount = .5
	}

	return originalPrice * discount
}
