package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getInteger(reader *bufio.Reader) (int, error) {
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return 0, err
	}

	number, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		return 0, err
	}

	return number, nil
}

func getFloat(reader *bufio.Reader) (float64, error) {
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return 0, err
	}

	number, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	amountOfStudents, err := getInteger(reader)
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}

	for range amountOfStudents {
		score, err := getFloat(reader)
		if err != nil {
			fmt.Printf("ошибка: %s\n", err)
			return
		}

		switch {
		case score >= 90 && score <= 100:
			fmt.Println(5)
		case score >= 75 && score <= 89:
			fmt.Println(4)
		case score >= 50 && score <= 74:
			fmt.Println(3)
		case score >= 0 && score <= 49:
			fmt.Println(2)
		default:
			fmt.Println("Неверный балл")
		}
	}
}
