package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	numberOfPurchases, err := getInteger(reader)
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}

	discount, err := getInteger(reader)
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}
	discountСoefficient := (100 - float64(discount)) / 100

	var price float64 = 0
	for range numberOfPurchases {
		purchasePrice, err := getFloat(reader)
		if err != nil {
			fmt.Printf("ошибка: %s\n", err)
			return
		}

		price += purchasePrice
		price = math.Round(price*1000) / 1000
	}

	fmt.Println(price * discountСoefficient)
}
