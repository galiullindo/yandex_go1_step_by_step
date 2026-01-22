package main

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	area := math.Pi * math.Pow(c.radius, 2)
	return math.Round(area*math.Pow(10, 14)) / math.Pow(10, 14)
}

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Area() float64 {
	area := r.width * r.height
	return math.Round(area*math.Pow(10, 14)) / math.Pow(10, 14)
}
