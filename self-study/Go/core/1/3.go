package main

import (
	"fmt"
)

// Запросить у пользователя температуру в градусах Цельсия и перевести её в градусы Фаренгейта и Кельвина.

func main () {
	var celsius float64
	
	fmt.Print("Enter temperature in Celsius: ")
	fmt.Scan(&celsius)
	
	// Фаренгейт = (Цельсий * 9/5) + 32
	fahrenheit := (celsius * 9 / 5) + 32

	// Кельвин = Цельсий + 273.15
	kelvin := celsius + 273.15

	fmt.Printf("Temperature in Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Temperature in Kelvin: %.2f\n", kelvin)
}