package main

import (
	"sort"
)

func SortAndMerge(s1 []int, s2 []int) []int {
	length := len(s1) + len(s2)
	merge := make([]int, length)

	sort.Ints(s1)
	sort.Ints(s2)

	i1 := 0
	i2 := 0
	for i := range length {
		switch {
		case i1 < len(s1) && i2 < len(s2):
			if s1[i1] <= s2[i2] {
				merge[i] = s1[i1]
				i1++
			} else {
				merge[i] = s2[i2]
				i2++
			}
		case i1 < len(s1):
			merge[i] = s1[i1]
			i1++
		case i2 < len(s2):
			merge[i] = s2[i2]
			i2++
		}
	}
	return merge
}
