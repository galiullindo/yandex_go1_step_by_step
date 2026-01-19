package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

	text, err := getText(reader)
	if err != nil {
		fmt.Println("Ошибка: некорректный ввод")
		return
	}

	data := strings.Split(text, " ")
	if len(data) != 3 {
		fmt.Println("Ошибка: некорректный ввод")
		return
	}

	age, err := strconv.Atoi(data[0])
	if err != nil {
		fmt.Println(err)
	}
	if age < 14 || age > 150 {
		fmt.Println("Ошибка: невалидный возраст")
		return
	}

	name := data[1]
	if utf8.RuneCountInString(name) < 2 {
		fmt.Println("Ошибка: невалидное имя")
		return
	}

	passportSeriesAndNumber := data[2]
	if utf8.RuneCountInString(passportSeriesAndNumber) != 10 {
		fmt.Println("Ошибка: невалидная серия и номер паспорта")
		return
	}

	fmt.Printf("Имя: %s. Возраст: %d. Серия и номер паспорта: %s\n", name, age, passportSeriesAndNumber)
}
