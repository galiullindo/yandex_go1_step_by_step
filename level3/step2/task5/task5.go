package main

func findMax(array [10]int) int {
	maxNumber := array[0]
	for _, number := range array {
		maxNumber = max(maxNumber, number)
	}
	return maxNumber
}

func findMin(array [10]int) int {
	minNumber := array[0]
	for _, number := range array {
		minNumber = min(minNumber, number)
	}
	return minNumber
}

func FindMaxMinInArray(array [10]int) (int, int) {
	maxNumber := findMax(array)
	minNumber := findMin(array)
	return maxNumber, minNumber
}
