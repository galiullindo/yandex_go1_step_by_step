package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getInteger(r *bufio.Reader) (int, error) {
	line, err := r.ReadString('\n')

	if err != nil && err != io.EOF {
		return 0, err
	}

	number, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		return 0, err
	}

	return number, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	number, err := getInteger(reader)
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}

	if number == 0 {
		fmt.Println("Число 0")
		return
	}
	if number >= -9 && number <= 9 {
		fmt.Println("Число однозначное")
		return
	}
	if number%2 == 0 {
		fmt.Println("Число чётное")
		return
	}
	if number > 0 {
		fmt.Println("Число положительное")
		return
	}

	fmt.Println("Число красивое")
}
