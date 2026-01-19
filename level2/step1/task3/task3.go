package main

import "fmt"

func GoOrNot(s string) {
	switch s {
	case "Go":
		fmt.Println("Go!")
	default:
		fmt.Println("Я знаю только Go!")
	}
}
