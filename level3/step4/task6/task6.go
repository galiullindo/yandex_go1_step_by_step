package main

func SumOfValuesInMap(m map[int]int) int {
	sum := 0
	for _, value := range m {
		sum += value
	}
	return sum
}
