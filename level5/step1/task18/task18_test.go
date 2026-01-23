package main

import "testing"

type test struct {
	name     string
	text     string
	expected string
}

var tests = []test{
	{
		"Lower case letters",
		"abcdefghijklmnopqrstuvwxyz",
		"bcdfghjklmnpqrstvwxyz",
	},
	{
		"Upper case letters",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"BCDFGHJKLMNPQRSTVWXYZ",
	},
	{
		"Repeating lower case letters",
		"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
		"bcdfghjklmnpqrstvwxyzbcdfghjklmnpqrstvwxyzbcdfghjklmnpqrstvwxyz",
	},
	{
		"Repeating upper case letters",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"BCDFGHJKLMNPQRSTVWXYZBCDFGHJKLMNPQRSTVWXYZBCDFGHJKLMNPQRSTVWXYZ",
	},
}

func TestDeleteVowels(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := DeleteVowels(test.text); got != test.expected {
				t.Errorf("DeleteVowels(\"%s\"), got \"%s\", expected \"%s\"\n", test.text, got, test.expected)
			}
		})
	}
}
