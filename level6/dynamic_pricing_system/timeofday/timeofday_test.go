package timeofday

import (
	"testing"
	"time"
)

type test struct {
	name     string
	t        time.Time
	expected float64
}

var tests = []test{}

func generateWeekdays() {
	day := 19 // 2026-01-19 - Понедельник
	times := 5
	for offset := range times {
		tests = append(
			tests,
			test{"Weekdays, night tariff", time.Date(2026, time.January, day+offset, 0, 0, 0, 0, time.UTC), 1.5},
			test{"Weekdays, night tariff", time.Date(2026, time.January, day+offset, 4, 59, 59, 1e9-1, time.UTC), 1.5},
		)
	}
	for offset := range times {
		tests = append(
			tests,
			test{"Weekdays, base tariff", time.Date(2026, time.January, day+offset, 5, 0, 0, 0, time.UTC), 1.0},
			test{"Weekdays, base tariff", time.Date(2026, time.January, day+offset, 6, 59, 59, 1e9-1, time.UTC), 1.0},
			test{"Weekdays, base tariff", time.Date(2026, time.January, day+offset, 10, 0, 0, 0, time.UTC), 1.0},
			test{"Weekdays, base tariff", time.Date(2026, time.January, day+offset, 23, 59, 59, 1e9-1, time.UTC), 1.0},
		)
	}
	for offset := range times {
		tests = append(
			tests,
			test{"Weekdays, morning rush hour", time.Date(2026, time.January, day+offset, 7, 0, 0, 0, time.UTC), 1.3},
			test{"Weekdays, morning rush hour", time.Date(2026, time.January, day+offset, 9, 59, 59, 1e9-1, time.UTC), 1.3},
		)
	}
}

func generateWeekend() {
	day := 24 // 2026-01-24 - Суббота
	times := 2
	for offset := range times {
		tests = append(
			tests,
			test{"Weekend, night tariff", time.Date(2026, time.January, day+offset, 0, 0, 0, 0, time.UTC), 1.5},
			test{"Weekend, night tariff", time.Date(2026, time.January, day+offset, 4, 59, 59, 1e9-1, time.UTC), 1.5},
		)
	}
	for offset := range times {
		tests = append(
			tests,
			test{"Weekend, weekend tariff", time.Date(2026, time.January, day+offset, 5, 0, 0, 0, time.UTC), 1.2},
			test{"Weekend, weekend tariff", time.Date(2026, time.January, day+offset, 6, 59, 59, 1e9-1, time.UTC), 1.2},
			test{"Weekend, weekend tariff", time.Date(2026, time.January, day+offset, 7, 0, 0, 0, time.UTC), 1.2},
			test{"Weekend, weekend tariff", time.Date(2026, time.January, day+offset, 9, 59, 59, 1e9-1, time.UTC), 1.2},
			test{"Weekend, weekend tariff", time.Date(2026, time.January, day+offset, 10, 0, 0, 0, time.UTC), 1.2},
			test{"Weekend, weekend tariff", time.Date(2026, time.January, day+offset, 23, 59, 59, 1e9-1, time.UTC), 1.2},
		)
	}
}

func TestGetTimeMultiplier(t *testing.T) {
	generateWeekdays()
	generateWeekend()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := GetTimeMultiplier(test.t); got != test.expected {
				t.Errorf("GetTimeMultiplier(%s) got %0.3f, expected %0.3f\n", test.t, got, test.expected)
			}
		})
	}
}
