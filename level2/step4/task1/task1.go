package main

import (
	"errors"
	"strings"
	"unicode/utf8"
	// "time" // для загрузки в Yandex-LMS необходимо добавлять импорт пакета иначе решение не работает.
)

var ErrAnswerIsIncorrect = errors.New("исправь свой ответ, а лучше ложись поспать")
var ErrInvalidDayOfWeek = errors.New("неверный день недели")

func currentDayOfTheWeek() string {
	now := TimeNow()
	dayOfWeek := now.Weekday()

	switch dayOfWeek {
	case 0:
		return "Воскресенье"
	case 1:
		return "Понедельник"
	case 2:
		return "Вторник"
	case 3:
		return "Среда"
	case 4:
		return "Четверг"
	case 5:
		return "Пятница"
	case 6:
		return "Суббота"
	default:
		panic(ErrInvalidDayOfWeek)
	}
}

func dayOrNight() string {
	now := TimeNow()
	hour := now.Hour()

	if hour > 9 && hour < 22 {
		return "День"
	}

	return "Ночь"
}

func nextFriday() int {
	now := TimeNow()
	dayOfWeek := now.Weekday()

	switch dayOfWeek {
	case 5:
		return 0
	case 6:
		return 7
	default:
		return 5 - int(dayOfWeek)
	}
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	dayOfWeek := currentDayOfTheWeek()
	if strings.EqualFold(answer, dayOfWeek) {
		return true
	}

	return false
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, ErrAnswerIsIncorrect
	}

	partOfDay := dayOrNight()
	if strings.EqualFold(answer, partOfDay) {
		return true, nil
	}

	return false, nil
}
