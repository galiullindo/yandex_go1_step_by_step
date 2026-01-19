package main

import (
	"unicode"
	"unicode/utf8"
)

func hasMinimumLength(passwd string, minimumLength int) bool {
	if utf8.RuneCountInString(passwd) < minimumLength {
		return false
	}

	return true
}

func hasUpper(passwd string) bool {
	for _, letter := range passwd {
		if unicode.IsUpper(letter) {
			return true
		}
	}

	return false
}

func hasLowerCase(passwd string) bool {
	for _, letter := range passwd {
		if unicode.IsLower(letter) {
			return true
		}
	}

	return false
}

func ratePassword(passwd string) string {
	var rate int
	var message string

	if hasMinimumLength(passwd, 8) {
		rate++
	}
	if hasUpper(passwd) {
		rate++
	}
	if hasLowerCase(passwd) {
		rate++
	}

	switch rate {
	case 0:
		message = "Пароль недостаточно безопасен, придумайте новый"
	case 1:
		message = "Слабый пароль"
	case 2:
		message = "Средний пароль"
	case 3:
		message = "Сложный пароль"
	}

	return message
}
