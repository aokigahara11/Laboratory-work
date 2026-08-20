package main

import (
	"fmt"
	"math"
)

// Определить, попадает ли точка с координатами $(x, y)$ в заштрихованную область (квадрант или круг заданного радиуса).

func main() {
	var x, y, r float64

	fmt.Print("Enter coordinates of point (x, y): ")
	fmt.Scan(&x, &y)

	fmt.Println("Enter radius of circle: ")
	fmt.Scan(&r)

	if x >= 0 && y >= 0 && math.Hypot(x, y) <= r {
    	fmt.Println("Point is in the shaded area")
	} else {
    	fmt.Println("Point is not in the shaded area")
	}

}
