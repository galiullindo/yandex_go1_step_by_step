package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	reader := bufio.NewReader(os.Stdin)

	firstPlayer, err := getText(reader)
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}

	secondPlayer, err := getText(reader)
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}

	switch {
	case firstPlayer == secondPlayer:
		fmt.Println("Ничья")
	case firstPlayer == "камень" && secondPlayer == "ножницы", firstPlayer == "ножницы" && secondPlayer == "бумага", firstPlayer == "бумага" && secondPlayer == "камень":
		fmt.Println("Первый игрок победил")
	case secondPlayer == "камень" && firstPlayer == "ножницы", secondPlayer == "ножницы" && firstPlayer == "бумага", secondPlayer == "бумага" && firstPlayer == "камень":
		fmt.Println("Второй игрок победил")
	default:
		fmt.Printf("паника: %s, %s\n", firstPlayer, secondPlayer)
	}
}
