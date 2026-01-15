package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getFloats(reader *bufio.Reader) ([]float64, error) {
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}

	numbersInStrings := strings.Split(strings.TrimSpace(line), " ")
	numbers := []float64{}
	for _, numberInString := range numbersInStrings {
		number, err := strconv.ParseFloat(numberInString, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	numbers, err := getFloats(reader)
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}
	if len(numbers) != 3 {
		fmt.Printf("ошибка: в строке не 3 числа")
		return
	}

	first, second, third := numbers[0], numbers[1], numbers[2]
	if first == second && second == third {
		fmt.Println("Максимальное равенство")
		return
	}

	fmt.Println("Не равны")
}
