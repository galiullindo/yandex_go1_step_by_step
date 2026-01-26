package main

import "sort"

func SortNums(nums []uint) {
	sort.Slice(nums, func(i int, j int) bool { return nums[i] < nums[j] })
}
