package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getText(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	return strings.TrimSpace(line), nil
}

func main() {
	var temperature int
	reader := bufio.NewReader(os.Stdin)

	text, err := getText(reader)
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}

	if text == "0" {
		temperature = 0
	} else {
		parts := strings.Split(text, " ")
		if len(parts) != 2 {
			fmt.Printf("ошибка: в строке не 2 числа")
			return
		}

		sign, numberInString := parts[0], parts[1]
		temperature, err = strconv.Atoi(numberInString)
		if err != nil {
			fmt.Printf("ошибка: %s\n", err)
			return
		}

		if sign == "-" {
			temperature = -1 * temperature
		}
	}

	var message string
	if temperature > 20 {
		message = "Стоит надеть майку и шорты"
	} else if temperature >= 10 && temperature <= 20 {
		message = "Стоит надеть штаны и кофту"
	} else if temperature >= -5 && temperature <= 9 {
		message = "Стоит надеть куртку"
	} else {
		message = "Стоит надеть зимнюю куртку"
	}

	fmt.Println(message)
}
