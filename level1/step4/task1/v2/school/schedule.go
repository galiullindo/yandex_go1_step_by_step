package school

import (
	"fmt"
)

type Schedule struct {
	minSlot       int
	maxSlot       int
	freeSlots     int
	occupiedSlots int
	slots         []*Slot
}

func NewShedule(size int) (*Schedule, error) {
	s := Schedule{minSlot: 1, maxSlot: size, freeSlots: size, occupiedSlots: 0, slots: []*Slot{}}
	for numbern := range size {
		slot, err := NewSlot(numbern + 1)
		if err != nil {
			return nil, err
		}
		s.slots = append(s.slots, slot)
	}
	return &s, nil
}

func (s *Schedule) ViewSchedule() string {
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

func (s *Schedule) ViewSlots() string {
	return fmt.Sprintf("Осталось свободных мест: %d\nВсего человек в очереди: %d\n", s.freeSlots, s.occupiedSlots)
}

func (s *Schedule) SignUp(number int, student string) error {
	if number < s.minSlot || number > s.maxSlot {
		return fmt.Errorf("Запись на место номер %d невозможна: некорректный ввод", number)
	}

	if s.freeSlots == 0 {
		return fmt.Errorf("Запись на место номер %d невозможна: очередь переполнена", number)
	}

	for _, slot := range s.slots {
		if slot.number == number {
			err := slot.SetStudent(student)
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
