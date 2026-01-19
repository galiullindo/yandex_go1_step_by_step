package school

import (
	"fmt"
	"testing"
)

type newSlotTestData struct {
	number      int
	wantErr     error
	wantNumber  int
	wantStudent string
}

func TestNewSlot(t *testing.T) {
	tests := []newSlotTestData{
		{number: -1, wantErr: ErrLessThanOrEqualZero, wantNumber: -1, wantStudent: ""},
		{number: 0, wantErr: ErrLessThanOrEqualZero, wantNumber: 0, wantStudent: ""},
		{number: 1, wantErr: nil, wantNumber: 1, wantStudent: ""},
		{number: 2, wantErr: nil, wantNumber: 2, wantStudent: ""},
	}

	for _, test := range tests {
		slot, gotErr := NewSlot(test.number)

		if gotErr != test.wantErr {
			t.Errorf("slot, err := NewSlot(%v), err = %s, want %s\n", test.number, gotErr, test.wantErr)
		}

		if slot != nil {
			if gotNumber := slot.Number(); gotNumber != test.wantNumber {
				t.Errorf("NewSlot(%v), slot.Number() = %d, want %d\n", test.number, gotNumber, test.wantNumber)
			}
			if gotStudent := slot.Student(); gotStudent != test.wantStudent {
				t.Errorf("NewSlot(%v), slot.Number() = %s, want %s\n", test.number, gotStudent, test.wantStudent)
			}
		}
	}
}

func TestSetStudent(t *testing.T) {
	var wantErr error

	slot, err := NewSlot(1)
	if err != nil {
		t.Error(err)
	}

	wantErr = nil
	if gotErr := slot.SetStudent("test"); gotErr != wantErr {
		t.Errorf("err := slot.SetStudent(\"test\"), err = %s, want %s\n", gotErr, wantErr)
	}

	wantErr = fmt.Errorf("Запись на место номер %d невозможна: место уже занято", slot.number)
	if gotErr := slot.SetStudent("test"); gotErr.Error() != wantErr.Error() {
		t.Errorf("err := slot.SetStudent(\"test\"), err = %s, want %s\n", gotErr, wantErr)
	}
}
