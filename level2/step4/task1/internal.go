package main

import (
	"fmt"
	"time"
)

func TimeNow() time.Time {
	layout := "2006-01-02 15:04:05 -0700 MST"
	now, _ := time.Parse(layout, "2025-06-03 12:15:33 +0000 UTC")
	return now
}

func testCurrentDayOfTheWeek() {
	want := "Вторник"
	got := currentDayOfTheWeek()

	if got != want {
		fmt.Printf("current day of the week: fail: %s != %s\n", got, want)
	} else {
		fmt.Printf("current day of the week: pass: %s == %s\n", got, want)
	}
}

func testDayOrNight() {
	want := "День"
	got := dayOrNight()

	if got != want {
		fmt.Printf("day or night: fail: %s != %s\n", got, want)
	} else {
		fmt.Printf("day or night: pass: %s == %s\n", got, want)
	}
}

func testNextFriday() {
	want := 3
	got := nextFriday()

	if got != want {
		fmt.Printf("next friday: fail: %d != %d\n", got, want)
	} else {
		fmt.Printf("next friday: pass: %d == %d\n", got, want)
	}
}

func main() {
	testCurrentDayOfTheWeek()
	testDayOrNight()
	testNextFriday()
}
