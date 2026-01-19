package main

import "fmt"

func PrintComplexNumber(z complex64) {
	fmt.Printf("Действительная часть: %0.2f. Мнимая часть: %0.2f\n", real(z), imag(z))
}
