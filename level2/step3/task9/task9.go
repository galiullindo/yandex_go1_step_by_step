package main

import (
	"fmt"
	"strings"
)

func CheckLetters(text string) string {
	eCounter := strings.Count(text, "е")
	if eCounter == 0 {
		return "Текст готов к публикации!"
	}

	return fmt.Sprintf("Количество возможных ошибок: %d, перепроверьте текст", eCounter)
}
