package main

import "strings"

func toLetters(r rune) string {
	letters := string(r)

	switch r {
	case '0':
		letters = "ноль"
	case '1':
		letters = "один"
	case '2':
		letters = "два"
	case '3':
		letters = "три"
	case '4':
		letters = "четыре"
	case '5':
		letters = "пять"
	case '6':
		letters = "шесть"
	case '7':
		letters = "семь"
	case '8':
		letters = "восемь"
	case '9':
		letters = "девять"
	case '+':
		letters = "плюс"
	case '-':
		letters = "минус"
	case '*':
		letters = "умножить на"
	case '/':
		letters = "разделить на"
	}

	return letters
}

func NumbersToLetters(input string) string {
	builder := strings.Builder{}
	for _, r := range []rune(input) {
		builder.WriteString(toLetters(r))
	}

	return builder.String()
}
