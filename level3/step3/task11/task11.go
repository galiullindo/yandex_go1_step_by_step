package main

import "errors"

var (
	ErrInvalidN = errors.New("invalid value of n")
)

func filter(prices []int, limit int) []int {
	filtred := make([]int, 0)
	for _, price := range prices {
		if price < limit {
			filtred = append(filtred, price)
		}
	}
	return filtred
}

func UnderLimit(prices []int, limit int, n int) ([]int, error) {
	if n <= 0 {
		return nil, ErrInvalidN
	}

	filtred := filter(prices, limit)
	if len(filtred) > n {
		return filtred[:n], nil
	}

	return filtred, nil
}
