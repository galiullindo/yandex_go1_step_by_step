package main

import "testing"

type test struct {
	a        int
	b        int
	expected int
}

var tests = []test{
	{0, 0, 0},
	{0, 1, 0},
	{0, 2, 0},
	{1, 2, 2},
	{2, 3, 6},
}

func TestMultiply(t *testing.T) {
	for _, test := range tests {
		t.Run("Case A multiply by B", func(t *testing.T) {
			if got := Multiply(test.a, test.b); got != test.expected {
				t.Errorf("Multiply(%d, %d), got %d, expected %d\n", test.a, test.b, got, test.expected)
			}
		})
		t.Run("Case B multiply by A", func(t *testing.T) {
			if got := Multiply(test.b, test.a); got != test.expected {
				t.Errorf("Multiply(%d, %d), got %d, expected %d\n", test.b, test.a, got, test.expected)
			}
		})
	}
}
