package main

import "testing"

type test struct {
	name     string
	a        int
	b        int
	expected int
}

var tests = []test{
	{"Zero and zero", 0, 0, 0},
	{"Zero and positive", 0, 1, 1},
	{"Zero and positive", 1, 0, 1},
	{"Zero and negative", 0, -1, -1},
	{"Zero and negative", -1, 0, -1},
	{"Positive and positive", 1, 1, 2},
	{"Positive and positive", 1, 1, 2},
	{"Positive and negative", 1, -1, 0},
	{"Positive and negative", -1, 1, 0},
	{"Negative and negative", -1, -1, -2},
	{"Negative and negative", -1, -1, -2},
}

func TestSum(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Sum(test.a, test.b)
			if got != test.expected {
				t.Errorf("Sum(%d, %d), got %d, expected %d\n", test.a, test.b, got, test.expected)
			}
		})
	}
}
