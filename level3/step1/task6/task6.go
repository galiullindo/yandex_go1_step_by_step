package main

import "unicode/utf8"

func CheckOnlyASCII(s string) bool {
	for _, r := range []rune(s) {
		if utf8.RuneLen(r) != 1 {
			return false
		}
	}
	return true
}
