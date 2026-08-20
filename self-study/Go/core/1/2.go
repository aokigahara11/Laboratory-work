package main

import (
	"fmt"
	"math"
)

// Даны катеты прямоугольного треугольника a и b. Найти его гипотенузу и площадь.

func main() {
	var a,b float64
	fmt.Print("Enter the lengths of the two legs of the right triangle: ")
	fmt.Scan(&a, &b)

	hypotenuse := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	area := (a * b) / 2

	fmt.Printf("Hypotenuse: %.2f\n", hypotenuse)
	fmt.Printf("Area: %.2f\n", area)
}