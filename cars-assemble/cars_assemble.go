package cars

const (
	oneHundredPercentage = float64(100)
	minutesInHour        = 60
	oneCost              = 10_000
	tenCost              = 95_000
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / oneHundredPercentage
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / minutesInHour)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const ten = 10
	return uint((carsCount/ten)*tenCost + (carsCount%ten)*oneCost)
}
