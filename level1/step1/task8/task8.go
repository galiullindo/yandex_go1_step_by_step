package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const kopecksInRuble = 100

func getIntegers(reader *bufio.Reader) ([]int, error) {
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}

	numbersInStrings := strings.Split(strings.TrimSpace(line), " ")
	numbers := []int{}
	for _, numberInString := range numbersInStrings {
		number, err := strconv.Atoi(numberInString)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	numbers, err := getIntegers(reader)
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}
	if len(numbers) != 3 {
		fmt.Printf("ошибка: в строке не 3 числа")
		return
	}

	rubles, kopecks, price := numbers[0], numbers[1], numbers[2]
	rubles += kopecks / kopecksInRuble
	kopecks = kopecks % kopecksInRuble
	if rubles >= price {
		fmt.Println("Сегодня будет вкусный кофе!")
		return
	}

	fmt.Println("Стоит подкопить")

}
