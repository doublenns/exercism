package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	var pick string
	if option1 < option2 {
		pick = option1
	} else {
		pick = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", pick)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var fairPrice float64
	switch {
	case age < 3:
		fairPrice = originalPrice * .8
	case age < 10:
		fairPrice = originalPrice * .7
	default:
		fairPrice = originalPrice * .5
	}
	return fairPrice
}
