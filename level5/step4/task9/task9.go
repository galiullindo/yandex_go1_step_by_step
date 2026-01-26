package main

import "sort"

func SortNames(names []string) {
	sort.Slice(names, func(i, j int) bool {
		nameI := []rune(names[i])
		nameJ := []rune(names[j])

		for i := range min(len(nameI), len(nameJ)) {
			if nameI[i] != nameJ[i] {
				return nameI[i] < nameJ[i]
			}
		}
		return false
	})
}
