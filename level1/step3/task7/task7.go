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

	for range 5 {
		text, err := getText(reader)
		if err != nil {
			fmt.Printf("ошибка: %s\n", err)
			return
		}

		if text == "Go" {
			fmt.Println("Go!")
		} else {
			fmt.Println("Я знаю только Go!")
		}
	}
}
