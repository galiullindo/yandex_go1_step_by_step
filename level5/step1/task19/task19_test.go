package main

import (
	"testing"
)

type test struct {
	name     string
	bytes    []byte
	expected int
	err      error
}

var tests = []test{
	{"Error return testing", []byte{0xff}, 0, ErrInvalidUTF8},
	{"Error return testing", []byte{0xfe}, 0, ErrInvalidUTF8},
	{"Error return testing", []byte{0xc3, 0x28}, 0, ErrInvalidUTF8},
	{"Error return testing", []byte{0x10, 0xe2, 0x28, 0x10}, 0, ErrInvalidUTF8},
	// ... and more invalid utf8 bytes ...

	{
		"UTF length return testing",
		[]byte{},
		0,
		nil,
	},
	{
		"Error return testing",
		[]byte{0},
		1,
		nil,
	},
	{
		"UTF length return testing",
		[]byte("abcdefghijklmnopqrstuvwxyz"),
		26,
		nil,
	},
	{
		"UTF length return testing",
		[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"),
		52,
		nil,
	},
}

func TestGetUTFLength(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got, err := GetUTFLength(test.bytes); got != test.expected || err != test.err {
				t.Errorf("GetUTFLength(%#v), got %d, %s, expected %d, %s\n", test.bytes, got, err, test.expected, test.err)
			}
		})
	}
}
