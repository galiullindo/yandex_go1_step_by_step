package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func makeInvitation(name string, apartment int, password int, duration float64) string {
	message := ""
	message += fmt.Sprintf("Привет, %s! ", name)
	message += fmt.Sprintf("Приглашаю тебя на соревнование по программированию, которое пройдёт, как всегда, в квартире %d. ", apartment)
	message += fmt.Sprintf("Оно будет длиться примерно %.1f часа. ", duration)
	message += fmt.Sprintf("Не забудь секретный пароль для входа: %d.", password)

	return message
}

func main() {
	var line string
	reader := bufio.NewReader(os.Stdin)

	line, _ = reader.ReadString('\n')
	name := strings.ReplaceAll(line, "\n", "")

	line, _ = reader.ReadString('\n')
	apartment, _ := strconv.Atoi(strings.ReplaceAll(line, "\n", ""))

	line, _ = reader.ReadString('\n')
	password, _ := strconv.Atoi(strings.ReplaceAll(line, "\n", ""))

	line, _ = reader.ReadString('\n')
	duration, _ := strconv.ParseFloat(strings.ReplaceAll(line, "\n", ""), 64)

	message := makeInvitation(name, apartment, password, duration)
	fmt.Println(message)
}
