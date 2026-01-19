package main

import "fmt"

func PrintFlightRow(flightNumber string, fromCity string, toCity string, flightDuration float64, receptionDeskNumber int, gateNumber int, registrationHasEnded bool) {
	var message string

	if registrationHasEnded {
		message = fmt.Sprintf("| %s | %s—%s | регистрация закончилась, проходите к гейту: %d | длительность полёта %0.1f часа |\n", flightNumber, fromCity, toCity, gateNumber, flightDuration)
	} else {
		message = fmt.Sprintf("| %s | %s—%s | %d регистрация продолжается |\n", flightNumber, fromCity, toCity, receptionDeskNumber)
	}

	fmt.Print(message)
}
