package main

func Join(s1 []int, s2 []int) []int {
	c := cap(s1) + cap(s2)
	s := make([]int, 0, c)
	s = append(s, s1...)
	s = append(s, s2...)
	return s
}
