package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("ошибка: %s\n", err)
	}

	text := strings.TrimSpace(line)
	if text == "Go" {
		fmt.Println("Go!")
		return
	}

	fmt.Println("Я знаю только Go!")
}
