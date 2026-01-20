package main

func FiveSteps(array [5]int) [5]int {
	revers := [5]int{}
	for i, e := range array {
		revers[4-i] = e
	}
	return revers
}
