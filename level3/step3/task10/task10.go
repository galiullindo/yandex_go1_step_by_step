package main

func Mix(s []int) []int {
	l := len(s)
	new := make([]int, 0, l)
	for i := range l / 2 {
		new = append(new, s[i], s[l/2+i])
	}
	return new
}
