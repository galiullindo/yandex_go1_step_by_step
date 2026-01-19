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

func checkPassword(passwd string) bool {
	if !hasMinimumLength(passwd, 8) {
		return false
	}
	if !hasUpper(passwd) {
		return false
	}

	return true
}
