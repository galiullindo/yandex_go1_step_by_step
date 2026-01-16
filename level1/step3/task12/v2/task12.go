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

	replacer := strings.NewReplacer("а", "", "о", "")
	encryptedWord := replacer.Replace(word)

	fmt.Println(encryptedWord)
}
