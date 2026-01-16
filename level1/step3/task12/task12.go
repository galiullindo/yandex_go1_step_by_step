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
	word := strings.TrimSpace(line)

	var encryptedWord string
	for _, letter := range word {
		if letter == 'а' {
			continue
		}
		if letter == 'о' {
			continue
		}

		encryptedWord += string(letter)
	}

	fmt.Println(encryptedWord)
}
