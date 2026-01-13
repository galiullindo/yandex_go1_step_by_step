package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var name string
	fmt.Println("Введите своё имя:")

	// fmt.Scanln(&name) // fmt.Scan-ы - для пидорасов этим не возможно пользоваться

	// Используй bufio.Reader для чтения
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	message := fmt.Sprintf("Привет %s!", name)
	fmt.Println(message)
}
