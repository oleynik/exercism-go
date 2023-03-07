// Package weather implements Forecast function provides information about
// current weather condition in requested city.
package weather

// CurrentCondition variable stores information about passed condition in the Forecast
// function and uses there to construct returned information.
var CurrentCondition string

// CurrentLocation variable stores name of the City passed as a parameter into the Forecast
// function and uses there to construct returned information.
var CurrentLocation string

// Forecast function accepts city and weather condition as a parameters and returns
// information about weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
