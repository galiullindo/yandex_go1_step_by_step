package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	timeOfBirth := time.Date(2001, 5, 9, 0, 0, 0, 0, time.UTC)
	diff := now.Sub(timeOfBirth)
	fmt.Println(diff.Hours())

	// Формат времени.
	// Мастер-шаблон для форматирования времени "2006-01-02T15:04:05Z07:00"
	fmt.Println("Шаблон для форматирования времени:", time.Layout)
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}
