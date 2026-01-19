package main

import "math"

func SquareRoots(a float64, b float64, c float64) (float64, float64) {
	var x1, x2 float64

	d := findDiscriminant(a, b, c)

	switch {
	case d < 0:
		x1, x2 = 0, 0
	case d == 0:
		x := -b / (2 * a)
		x1, x2 = x, x
	case d > 0:
		x1 = (-b + math.Sqrt(d)) / (2 * a)
		x2 = (-b - math.Sqrt(d)) / (2 * a)
	}

	if x1 > x2 {
		return x2, x1
	}

	return x1, x2
}

func findDiscriminant(a float64, b float64, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}
