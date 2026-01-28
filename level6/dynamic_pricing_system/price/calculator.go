package price

import (
	"math"
	"time"

	"github.com/galiullindo/yandex_go1_step_by_step/level6/dynamic_pricing_system/timeofday"
	"github.com/galiullindo/yandex_go1_step_by_step/level6/dynamic_pricing_system/traffic"
	"github.com/galiullindo/yandex_go1_step_by_step/level6/dynamic_pricing_system/weathr"
)

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
