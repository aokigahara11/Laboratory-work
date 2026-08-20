package main

import (
	"fmt"
)

// Проверить, является ли введенный год високосным.

func main() {
	var year int

	fmt.Print("Enter year: ")
	fmt.Scan(&year)

	if year % 4 == 0 && (year % 100 != 0 || year % 400 == 0) {
		fmt.Printf("%d is a leap year", year)
	} else {
		fmt.Printf("%d is not a leap year", year)
	}
}