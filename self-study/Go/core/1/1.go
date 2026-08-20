package main

import (
	"fmt"
	"math"
)

// Вычислить площадь и длину окружности по заданному радиусу R.

func main() {
	var radius float64
	const Pi = math.Pi
	
	fmt.Print("Enter radius of circle: ")
	fmt.Scan(&radius)

	area := Pi * math.Pow(radius, 2)
	circumference := 2 * Pi * radius

	fmt.Printf("Area: %.2f\n", area)
	fmt.Printf("Length of circle: %.2f\n", circumference)

}