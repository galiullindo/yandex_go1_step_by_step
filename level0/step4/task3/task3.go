package main

import (
	"fmt"
	"math"
)

func main() {
	// b := int64(2.1) // cannot convert 2.1 (untyped float constant) to type int64

	// Математика
	fmt.Println(math.Abs(-1))
	fmt.Println(math.Min(-1, 1))
	fmt.Println(math.Max(-1, 1))
	fmt.Println(math.Sqrt(2))
	fmt.Println(math.Sin(180))
	fmt.Println(math.Pow(2, 3))

	// Округление
	fmt.Println(math.Ceil(5.4))
	fmt.Println(math.Floor(5.6))
	fmt.Println(math.Round(5.4))
	fmt.Println(math.Round(5.5))
	fmt.Println(math.Round(5.6))
	fmt.Println(5.4, math.RoundToEven(5.4))
	fmt.Println(5.5, math.RoundToEven(5.5))
	fmt.Println(5.6, math.RoundToEven(5.6))
	fmt.Println(6.4, math.RoundToEven(6.4))
	fmt.Println(6.5, math.RoundToEven(6.5))
	fmt.Println(6.6, math.RoundToEven(6.6))

}
