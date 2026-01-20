package main

import "fmt"

func PrettyArrayOutput(tasks [9]string) {
	for index, task := range tasks {
		if index < 7 {
			fmt.Printf("%d я уже сделал: %s\n", index+1, task)
		} else {
			fmt.Printf("%d не успел сделать: %s\n", index+1, task)
		}
	}
}
