package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"yandex_lms/level1/step4/task1/v2/school"
)

func getEvent(r *bufio.Reader) (string, error) {
	e, err := r.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSpace(e), nil
}

func main() {
	var isRun bool = true
	reader := bufio.NewReader(os.Stdin)
	subjectSchedule, _ := school.NewShedule(5)

	for isRun {
		event, err := getEvent(reader)
		if err != nil {
			log.Fatal(err)
		}

		switch event {
		case "":
			continue
		case "конец":
			msg := subjectSchedule.ViewSchedule()
			fmt.Print(msg)
			isRun = false
			continue
		case "очередь":
			msg := subjectSchedule.ViewSchedule()
			fmt.Print(msg)
			continue
		case "количество":
			msg := subjectSchedule.ViewSlots()
			fmt.Print(msg)
			continue
		default:
			data := strings.Split(event, " ")
			if len(data) != 2 {
				log.Fatal("data length!= 2\n")
			}

			student := data[0]
			number, err := strconv.Atoi(data[1])
			if err != nil {
				fmt.Println(err)
			}

			err = subjectSchedule.SignUp(number, student)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}
