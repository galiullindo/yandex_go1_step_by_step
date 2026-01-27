package weathr

type WeatherCondition int

const (
	Clear WeatherCondition = iota
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	multiplier := 1.0

	switch weather.Condition {
	case HeavyRain:
		multiplier += 0.2 // Увеличиваем коэффициент в сильный дождь
	case Snow:
		multiplier += 0.15 // Увеличиваем коэффициент в снег
	case Rain:
		multiplier += 0.125 // Увеличиваем коэффициент в дождь
	}

	if weather.WindSpeed > 15 {
		multiplier += 0.1 // Увеличиваем коэффициент при сильном ветре
	}

	return multiplier
}
