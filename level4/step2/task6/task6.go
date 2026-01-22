package main

import "fmt"

type Animal interface {
	MakeSound()
}

type Cat struct {
	sound string
}

func (c Cat) MakeSound() {
	fmt.Println("Мяу!")
}

type Dog struct {
}

func (d Dog) MakeSound() {
	fmt.Println("Гав!")
}
