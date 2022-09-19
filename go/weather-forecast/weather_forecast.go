// Package weather forecasts the current weather condition of various
// cities in Goblinocus.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current city in Goblinocus.
var CurrentLocation string

// Forecast returns a string denoting the current weather conditions of a
// specific Goblinocus city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
