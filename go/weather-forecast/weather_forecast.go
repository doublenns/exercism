// Package weather forcasts the current weather condition of various cities
// in Goblinocus.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current city in Goblinocus.
var CurrentLocation string

// Forecast returns a string value describing the forecasted weather condition
// of a city in Goblinocus.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
