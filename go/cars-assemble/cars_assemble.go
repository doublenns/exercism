package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	const minutesInAnHour int = 60
	workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(workingCarsPerHour) / minutesInAnHour
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const individualCarCost uint = 10000
	const groupOfTenCarsCost uint = 95000
	cars := uint(carsCount)

	return ((cars / 10) * groupOfTenCarsCost) + ((cars % 10) * individualCarCost)
}
