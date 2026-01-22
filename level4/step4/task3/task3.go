package main

import (
	"math"
	"reflect"
)

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
}

type Car struct {
	Type     string
	Speed    float64
	FuelType string
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	travelTime := distance / c.Speed
	return math.Round(travelTime*math.Pow(10, 14)) / math.Pow(10, 14)
}

type Motorcycle struct {
	Type     string
	Speed    float64
	FuelType string
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	travelTime := distance / m.Speed
	return math.Round(travelTime*math.Pow(10, 14)) / math.Pow(10, 14)
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	mapTravelTimeByType := make(map[string]float64)
	for _, vehicle := range vehicles {
		travelTime := vehicle.CalculateTravelTime(distance)
		vehicleType := reflect.TypeOf(vehicle)
		mapTravelTimeByType[vehicleType.String()] = travelTime
	}
	return mapTravelTimeByType
}
