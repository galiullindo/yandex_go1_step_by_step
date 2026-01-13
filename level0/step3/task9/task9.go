package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Мне не нравится использовать fmt.Scan-ы по причине того,
	// что если ты вводишь в буфер данных больше чем переменных то остатки данных в буфере вылазят в консоль.
	/*
		var message string
		_, err := fmt.Scanln(&message)
		if err != nil {
			fmt.Printf("ошибка: %s\n", err)
		}
	*/

	// На мой взгляд это лучше.
	reader := bufio.NewReader(os.Stdin)
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
	}

	message = strings.ReplaceAll(message, "\n", "")
	fmt.Printf("А ещё я люблю %s!\n", message)
}
