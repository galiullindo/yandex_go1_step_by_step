package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func getText(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	return strings.TrimSpace(line), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstPasswd, err := getText(reader)
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
	}

	secondPasswd, err := getText(reader)
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
	}

	if utf8.RuneCountInString(firstPasswd) >= 8 && utf8.RuneCountInString(secondPasswd) >= 8 {
		fmt.Println("Оба пароля надёжные")
		return
	}
	if utf8.RuneCountInString(firstPasswd) >= 8 {
		fmt.Println("Только первый пароль надёжный")
		return
	}
	if utf8.RuneCountInString(secondPasswd) >= 8 {
		fmt.Println("Только второй пароль надёжный")
		return
	}

	fmt.Println("Оба пароля ненадёжные")
}
