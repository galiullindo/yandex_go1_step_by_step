package main

import (
	"fmt"
	"time"

	"github.com/galiullindo/yandex_go1_step_by_step/level6/dynamic_pricing_system/price"
	"github.com/galiullindo/yandex_go1_step_by_step/level6/dynamic_pricing_system/traffic"
	"github.com/galiullindo/yandex_go1_step_by_step/level6/dynamic_pricing_system/weathr"
)

func main() {
	calculator := price.PriceCalculator{
		TrafficClient: &traffic.RealTrafficClient{}, // В продакшене используется настоящий клиент, а мы подключим структуру-заглушку для имитации его работы
	}

	price := calculator.CalculatePrice(
		price.TripParameters{Distance: 8.5, Duration: 20},
		time.Now(),
		weathr.WeatherData{Condition: weathr.HeavyRain, WindSpeed: 10},
		55.751244, 37.618423,
	)

	fmt.Printf("Ваша цена: %.0f руб.\n", price)
}
