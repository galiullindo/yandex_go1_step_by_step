package main

import "unicode/utf8"

func isLong(s string, threshold int) bool {
	if utf8.RuneCountInString(s) < threshold {
		return false
	}
	return true
}

func DeleteLongKeys(m map[string]int) map[string]int {
	for k := range m {
		if !isLong(k, 6) {
			delete(m, k)
		}
	}
	return m
}
