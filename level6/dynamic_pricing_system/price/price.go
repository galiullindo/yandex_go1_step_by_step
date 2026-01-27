package price

import (
	"math"
	"time"
	"yandex_lms/level6/dynamic_pricing_system/timeofday"
	"yandex_lms/level6/dynamic_pricing_system/traffic"
	"yandex_lms/level6/dynamic_pricing_system/weathr"
)

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

type PriceCalculator struct {
	TrafficClient traffic.TrafficClient
}

func (c *PriceCalculator) CalculatePrice(trip TripParameters, now time.Time, weather weathr.WeatherData, lat, lng float64) float64 {
	base := CalculateBasePrice(trip)
	timeMult := timeofday.GetTimeMultiplier(now)
	weatherMult := weathr.GetWeatherMultiplier(weather)
	trafficMult := traffic.GetTrafficMultiplier(c.TrafficClient.GetTrafficLevel(lat, lng))

	finalPrice := base * timeMult * weatherMult * trafficMult

	return ApplyPriceLimits(math.Round(finalPrice))
}
