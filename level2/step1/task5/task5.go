package main

import "fmt"

func printPrice(price int, name string) {
	fmt.Printf("%s будет стоить %d рублей\n", name, price)
}

func BuyFries(size string) {
	switch size {
	case "S":
		printPrice(49, "Картошка фри")
	case "M":
		printPrice(89, "Картошка фри")
	case "L":
		printPrice(99, "Картошка фри")
	default:
		fmt.Println("Некорректный размер")
	}
}

func BuyCola(size string) {
	switch size {
	case "S":
		printPrice(99, "Кола")
	case "M":
		printPrice(119, "Кола")
	case "L":
		printPrice(139, "Кола")
	default:
		fmt.Println("Некорректный размер")
	}
}
