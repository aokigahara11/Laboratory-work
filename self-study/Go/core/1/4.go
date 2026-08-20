package main

import (
	"fmt"
	"math"
)

// Вычислить расстояние между двумя точками с координатами $(x_1, y_1)$ и $(x_2, y_2)$.

func main() {
	var x1, y1, x2, y2 float64

	fmt.Print("Enter coordinates first point: ")
	fmt.Scan(&x1, &y1)

	fmt.Print("Enter coordinates second point: ")
	fmt.Scan(&x2, &y2)

	dx := x2 - x1
	dy := y2 - y1

	distance := math.Hypot(dx, dy)

	fmt.Printf("Distance between points: %.2f\n", distance)
}