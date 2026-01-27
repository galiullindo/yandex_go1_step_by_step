package price

const (
	pricePerKm     = 10.0
	pricePerMinute = 2.0
)

type TripParameters struct {
	Distance float64
	Duration float64
}

func CalculateBasePrice(trip TripParameters) float64 {
	return trip.Distance*pricePerKm + trip.Duration*pricePerMinute
}

const (
	minPrice = 99.0
	maxPrice = 20000.0
)

func ApplyPriceLimits(price float64) float64 {
	if price < minPrice {
		return minPrice
	}
	if price > maxPrice {
		return maxPrice
	}
	return price
}
