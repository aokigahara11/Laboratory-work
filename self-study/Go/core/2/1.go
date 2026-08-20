package main

import (
	"fmt"

)

// Даны три действительных числа. Найти максимальное и минимальное из них.

func main() {
	var a, b, c float64

	fmt.Print("Enter three float numbers: ")
	fmt.Scan(&a, &b, &c)

	max := max(a, b, c)
	min := min(a, b, c)

	fmt.Printf("Max: %.2f, Min: %.2f", max, min)
}