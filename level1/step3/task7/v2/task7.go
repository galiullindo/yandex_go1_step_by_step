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

	var err error = nil
	var line string
	for err != io.EOF {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf("ошибка: %s\n", err)
			return
		}

		text := strings.TrimSpace(line)
		if text == "Go" {
			fmt.Println("Go!")
		} else {
			fmt.Println("Я знаю только Go!")
		}
	}
}
