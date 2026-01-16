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
	var line string
	var err error
	reader := bufio.NewReader(os.Stdin)

	line, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}

	numberOfWords, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}

	line, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("ошибка: %s\n", err)
		return
	}
	targetWord := strings.ToLower(strings.TrimSpace(line))

	targetWordCounter := 0
	for i := range numberOfWords {
		if i == numberOfWords-1 {
			line, err = reader.ReadString('\n')
		} else {
			line, err = reader.ReadString(' ')
		}

		if err != nil && err != io.EOF {
			fmt.Printf("ошибка: %s\n", err)
			return
		}

		word := strings.ToLower(strings.TrimSpace(line))
		if word == targetWord {
			targetWordCounter++
		}
	}

	fmt.Println(targetWordCounter)
}
