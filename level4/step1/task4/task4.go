package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   float64
	bonus    float64
}

func (e *Employee) CalculateTotalSalary() {
	totalSalary := e.salary + e.bonus
	fmt.Printf("Employee: %s\n", e.name)
	fmt.Printf("Position: %s\n", e.position)
	fmt.Printf("Total Salary: %0.2f\n", totalSalary)
}
