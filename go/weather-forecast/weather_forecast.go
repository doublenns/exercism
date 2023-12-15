// Package weather provides weather forcasts for different cities.
package weather

// CurrentCondition rerpresents the current weather condition.
var CurrentCondition string
// CurrentLocation represents the city currently located in.
var CurrentLocation string


// Forecast returns a string value of the weather, given the current condition and city. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
