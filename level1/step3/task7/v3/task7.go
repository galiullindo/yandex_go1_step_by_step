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

	for {
		text, err := getText(reader)
		if err != nil {
			fmt.Printf("ошибка: %s\n", err)
			return
		}

		if text == "" {
			break
		}

		switch text {
		case "Go":
			fmt.Println("Go!")
		default:
			fmt.Println("Я знаю только Go!")
		}
	}
}
