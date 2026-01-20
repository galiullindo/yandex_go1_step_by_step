package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func countLength(s string) int {
	return utf8.RuneCountInString(s)
}

func countBytes(s string) int {
	return len(s)
}

func CountLengthAndBytes(first string, second string) string {
	lengthSum := countLength(first) + countLength(second)
	bytesSum := countBytes(first) + countBytes(second)
	return fmt.Sprintf("Объединённая строка: %s. Количество байт: %d. Количество символов: %d.", strings.Join([]string{first, second}, ""), bytesSum, lengthSum)
}
