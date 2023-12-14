package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	var sr float64
	if speed < 1 {
		sr = 0
	} else if speed >= 1 && speed <= 4 {
		sr = 1.0
	} else if speed <= 8 {
		sr = .9
	} else if speed <= 10 {
		sr = .77
	}
	return sr
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	const cph float64 = 221
	return SuccessRate(speed) * float64(speed) * cph
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed)) / 60
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	unlimited := CalculateProductionRatePerHour(speed)
	if unlimited < limit {
		return unlimited
	} else {
		return limit
	}
}
