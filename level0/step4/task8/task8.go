package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func getMessage(hours int, minutes int) string {
	return fmt.Sprintf("Текущее время %d часов, %d минут. Ты точно не забыл про важные дела на сегодня?", hours, minutes)
}

func check(err error) {
	if err != nil && err != io.EOF {
		fmt.Printf("ошибка: %s\n", err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	check(err)

	layout := "2006-01-02/15:04:05"
	dateTime, err := time.Parse(layout, strings.TrimSpace(line))
	check(err)

	hours := dateTime.Hour()
	minutes := dateTime.Minute()
	message := getMessage(hours, minutes)

	fmt.Println(message)
}
