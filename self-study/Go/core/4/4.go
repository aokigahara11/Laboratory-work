package main

import (
	"fmt"
	"log"
)

// Написать функцию деления двух чисел, которая использует defer для логирования 
// завершения работы функции и корректно обрабатывает ошибку деления на ноль 
// без аварийного завершения программы.

func Div(a, b float64) (result float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic in Div: %v", r)
		}
		log.Printf("Division function completed with result: %v, error: %v", result, err)
	}

	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}

	result = a / b
	return result, nil
}

func main() {
	var a, b float64
	fmt.Print("Enter two numbers (a and b): ")
	fmt.Scan(&a, &b)

	result, err := Div(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result of division: %f\n", result)
	}
}