package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("ошибка: %s\n", err)
		return
	}

	text := strings.TrimSpace(line)
	numbers := strings.Split(text, " ")
	if len(numbers) != 2 {
		fmt.Printf("ошибка: в строке не 2 числа")
		return
	}

	firstNumber, err := strconv.Atoi(numbers[0])
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}
	secondNumber, err := strconv.Atoi(numbers[1])
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}

	if firstNumber == secondNumber {
		fmt.Println("Числа равны")
	} else if firstNumber > secondNumber {
		fmt.Println("Первое число больше второго")
	} else if firstNumber < secondNumber {
		fmt.Println("Второе число больше первого")
	}
}
