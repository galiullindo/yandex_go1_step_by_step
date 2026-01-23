package main

import "testing"

type test struct {
	name     string
	values   []int
	expected string
}

var tests = []test{
	{
		"Less than zero",
		[]int{-10, -1},
		"negative",
	},
	{
		"Equels zero",
		[]int{0},
		"zero",
	},
	{
		"More than zero and less than ten",
		[]int{1, 9},
		"short",
	},
	{
		"Ten or more and less than one hundred",
		[]int{10, 99},
		"long",
	},
	{
		"One hundred or more",
		[]int{100, 101},
		"very long",
	},
}

func TestLength(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, value := range test.values {
				if got := Length(value); got != test.expected {
					t.Errorf("Length(%d), got \"%s\", expected \"%s\"\n", value, got, test.expected)
				}
			}
		})
	}
}
