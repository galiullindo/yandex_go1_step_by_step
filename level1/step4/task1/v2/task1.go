package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type slot struct {
	number  int
	student string
}

func newSlot(number int) *slot {
	s := slot{number: number, student: ""}
	return &s
}

func (s *slot) setStudent(student string) error {
	if s.student != "" {
		return fmt.Errorf("Запись на место номер %d невозможна: место уже занято", s.number)
	}
	s.student = student
	return nil
}

type schedule struct {
	minSlot       int
	maxSlot       int
	freeSlots     int
	occupiedSlots int
	slots         []*slot
}

func newShedule(size int) *schedule {
	s := schedule{minSlot: 1, maxSlot: size, freeSlots: size, occupiedSlots: 0, slots: []*slot{}}
	for slotNumbern := range size {
		s.slots = append(s.slots, newSlot(slotNumbern+1))
	}
	return &s
}

func (s *schedule) viewSchedule() string {
	var view string
	for _, slot := range s.slots {
		placeholder := "-"
		if slot.student != "" {
			placeholder = slot.student
		}
		viewPart := fmt.Sprintf("%d. %s\n", slot.number, placeholder)
		view += viewPart
	}
	return view
}

func (s *schedule) viewSlots() string {
	return fmt.Sprintf("Осталось свободных мест: %d\nВсего человек в очереди: %d\n", s.freeSlots, s.occupiedSlots)
}

func (s *schedule) signUp(number int, student string) error {
	if number < s.minSlot || number > s.maxSlot {
		return fmt.Errorf("Запись на место номер %d невозможна: некорректный ввод", number)
	}

	if s.freeSlots == 0 {
		return fmt.Errorf("Запись на место номер %d невозможна: очередь переполнена", number)
	}

	for _, slot := range s.slots {
		if slot.number == number {
			err := slot.setStudent(student)
			if err != nil {
				return err
			}
			s.freeSlots--
			s.occupiedSlots++
			break
		}
	}

	return nil
}

func getEvent(r *bufio.Reader) (string, error) {
	e, err := r.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSpace(e), nil
}

func main() {
	var isRun bool = true

	subjectSchedule := newShedule(5)
	reader := bufio.NewReader(os.Stdin)

	for isRun {
		event, err := getEvent(reader)
		if err != nil {
			log.Fatal(err)
		}

		switch event {
		case "":
			continue
		case "конец":
			msg := subjectSchedule.viewSchedule()
			fmt.Print(msg)
			isRun = false
			continue
		case "очередь":
			msg := subjectSchedule.viewSchedule()
			fmt.Print(msg)
			continue
		case "количество":
			msg := subjectSchedule.viewSlots()
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

			err = subjectSchedule.signUp(number, student)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}
