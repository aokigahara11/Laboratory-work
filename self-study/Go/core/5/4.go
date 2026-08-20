package main

import (
	"fmt"
)

// Объявить структуру Rectangle с полями Width и Height (float64). 
// Написать методы для этой структуры, вычисляющие площадь и периметр прямоугольника.

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Println("Area:", rect.Area())
	fmt.Println("Perimeter:", rect.Perimeter())
}