package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	ErrInvalidInput    = errors.New("Ошибка: некорректный ввод")
	ErrInvalidAge      = errors.New("Ошибка: невалидный возраст")
	ErrInvalidName     = errors.New("Ошибка: невалидное имя")
	ErrInvalidPassport = errors.New("Ошибка: невалидная серия и номер паспорта")
)

func getText(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	return strings.TrimSpace(line), nil
}

func parseData(s string) (age int, name string, passport string, err error) {
	data := strings.Split(s, " ")
	if len(data) != 3 {
		return 0, "", "", ErrInvalidInput
	}

	age, err = strconv.Atoi(data[0])
	if err != nil {
		return 0, "", "", err
	}

	name = data[1]
	passport = data[2]

	return age, name, passport, nil
}

func checkAge(age int) error {
	if age < 14 || age > 150 {
		return ErrInvalidAge
	}
	return nil
}

func checkName(name string) error {
	if utf8.RuneCountInString(name) < 2 {
		return ErrInvalidName
	}
	return nil
}

func checkPassport(passport string) error {
	if utf8.RuneCountInString(passport) != 10 {
		return ErrInvalidPassport
	}
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, err := getText(reader)
	if err != nil {
		fmt.Println(err)
	}

	age, name, passport, err := parseData(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := checkAge(age); err != nil {
		fmt.Println(err)
		return
	}
	if err := checkName(name); err != nil {
		fmt.Println(err)
		return
	}
	if err := checkPassport(passport); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Имя: %s. Возраст: %d. Серия и номер паспорта: %s\n", name, age, passport)
}
