/*
Package weather provides weather forecasts for specified cities.
*/
package weather

var (
	// CurrentCondition specifies the current condition within the specified city.
	CurrentCondition string
	// CurrentLocation specifies the current city.
	CurrentLocation string
)

// Forecast takes in a city and a condition and returns a forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
