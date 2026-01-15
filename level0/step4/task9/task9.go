package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func getDate(reader *bufio.Reader, layout string) (time.Time, error) {
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return time.Time{}, err
	}

	stringDate := strings.TrimSpace(line)
	date, err := time.Parse(layout, stringDate)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func main() {
	layout := "2006-01-02"
	reader := bufio.NewReader(os.Stdin)

	futureTime, err := getDate(reader, layout)
	if err != nil && err != io.EOF {
		fmt.Printf("ошибка: %s\n", err)
	}

	presentTime, err := getDate(reader, layout)
	if err != nil && err != io.EOF {
		fmt.Printf("ошибка: %s\n", err)
	}

	diff := futureTime.Year() - presentTime.Year()
	fmt.Printf("%d year ago\n", diff)
}
