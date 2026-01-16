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
	const (
		maxPlaces = 5
	)

	var (
		isRun bool = true

		freePlaces     int = 5
		occupiedPlaces int = 0

		firstPlace  string = "-"
		secondPlace string = "-"
		thirdPlace  string = "-"
		fourthPlace string = "-"
		fifthPlace  string = "-"
	)

	reader := bufio.NewReader(os.Stdin)

	for isRun {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf("error: %s\n", err)
			return
		}

		event := strings.TrimSpace(line)
		if event == "" {
			continue
		}

		switch event {
		case "конец":
			msg := fmt.Sprintf("1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n", firstPlace, secondPlace, thirdPlace, fourthPlace, fifthPlace)
			fmt.Print(msg)
			isRun = false
			continue
		case "очередь":
			msg := fmt.Sprintf("1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n", firstPlace, secondPlace, thirdPlace, fourthPlace, fifthPlace)
			fmt.Print(msg)
			continue
		case "количество":
			msg := fmt.Sprintf("Осталось свободных мест: %d\nВсего человек в очереди: %d\n", freePlaces, occupiedPlaces)
			fmt.Print(msg)
			continue
		}

		parts := strings.Split(event, " ")
		if len(parts) != 2 {
			fmt.Printf("error: len != 2\n")
			return
		}

		studentName := parts[0]
		placeNumber, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}

		if placeNumber < 1 || placeNumber > 5 {
			err := "некорректный ввод"
			msg := fmt.Sprintf("Запись на место номер %d невозможна: %s", placeNumber, err)
			fmt.Println(msg)
			continue
		}
		if occupiedPlaces == maxPlaces {
			err := "очередь переполнена"
			msg := fmt.Sprintf("Запись на место номер %d невозможна: %s", placeNumber, err)
			fmt.Println(msg)
			continue
		}

		switch placeNumber {
		case 1:
			if firstPlace != "-" {
				err := "место уже занято"
				msg := fmt.Sprintf("Запись на место номер %d невозможна: %s", placeNumber, err)
				fmt.Println(msg)
				continue
			}

			firstPlace = studentName
			freePlaces--
			occupiedPlaces++

		case 2:
			if secondPlace != "-" {
				err := "место уже занято"
				msg := fmt.Sprintf("Запись на место номер %d невозможна: %s", placeNumber, err)
				fmt.Println(msg)
				continue
			}

			secondPlace = studentName
			freePlaces--
			occupiedPlaces++

		case 3:
			if thirdPlace != "-" {
				err := "место уже занято"
				msg := fmt.Sprintf("Запись на место номер %d невозможна: %s", placeNumber, err)
				fmt.Println(msg)
				continue
			}

			thirdPlace = studentName
			freePlaces--
			occupiedPlaces++
		case 4:
			if fourthPlace != "-" {
				err := "место уже занято"
				msg := fmt.Sprintf("Запись на место номер %d невозможна: %s", placeNumber, err)
				fmt.Println(msg)
				continue
			}

			fourthPlace = studentName
			freePlaces--
			occupiedPlaces++

		case 5:
			if fifthPlace != "-" {
				err := "место уже занято"
				msg := fmt.Sprintf("Запись на место номер %d невозможна: %s", placeNumber, err)
				fmt.Println(msg)
			} else {
				fifthPlace = studentName
				freePlaces--
				occupiedPlaces++
			}
		}
	}
}
