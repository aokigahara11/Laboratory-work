package main 

import (
	"fmt"
	"math"
)

// Написать функцию, которая принимает три числа и возвращает среднее арифметическое и среднее геометрическое.

func getAverages(a, b, c float64) (float64, float64) {
	arithmetic_mean := (a + b + c) / 3
	geometric_mean := math.Cbrt(a * b * c)
	return arithmetic_mean, geometric_mean
}

func main() {
	var a, b, c float64

	fmt.Print("Enter three float numbers: ")
	fmt.Scan(&a, &b, &c)

	arithmetic_mean, geometric_mean := getAverages(a, b, c)
	fmt.Printf("Arithmetic mean: %.2f, Geometric mean: %.2f", arithmetic_mean, geometric_mean)
}