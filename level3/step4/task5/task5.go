package main

func FindMaxKey(m map[int]int) int {
	maxKey := 0
	hasInitKey := false

	for key := range m {
		if !hasInitKey {
			maxKey = key
			hasInitKey = true
			continue
		}
		maxKey = max(key, maxKey)
	}

	return maxKey
}
