package school

import (
	"errors"
	"fmt"
)

const minSlotNumber = 1

var ErrLessThanOrEqualZero = errors.New("the slot number cannot be less than or equal to 0")

type Slot struct {
	number  int
	student string
}

func NewSlot(number int) (*Slot, error) {
	if number < minSlotNumber {
		return nil, ErrLessThanOrEqualZero
	}
	s := Slot{number: number, student: ""}
	return &s, nil
}

func (s *Slot) Number() int {
	return s.number
}

func (s *Slot) Student() string {
	return s.student
}

func (s *Slot) SetStudent(student string) error {
	if s.student != "" {
		return fmt.Errorf("Запись на место номер %d невозможна: место уже занято", s.number)
	}
	s.student = student
	return nil
}
