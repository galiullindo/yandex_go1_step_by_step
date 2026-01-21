package main

func CountingSort(contacts []string) map[string]int {
	repetitions := make(map[string]int)
	for _, contact := range contacts {
		repetitions[contact]++
	}
	return repetitions
}
