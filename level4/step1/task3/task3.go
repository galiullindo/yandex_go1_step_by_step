package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (p *Person) PrettyPrint() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Address:", p.address)
}
