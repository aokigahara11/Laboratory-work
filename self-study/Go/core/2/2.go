package main

import (
	"fmt"
)

// Написать программу-калькулятор: пользователь вводит два числа и символ операции (+, -, *, /). 
// Выполнить действие с помощью switch и обработать деление на ноль.

func main() {
	var a, b float64
	var operation string

	fmt.Print("Enter two float numbers: ")
	fmt.Scan(&a, &b)

	fmt.Print("Enter operation (+, -, *, /): ")
	fmt.Scan(&operation)

	switch operation {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f", a, b, a + b)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f", a, b, a - b)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f", a, b, a * b)
	case "/":
		if b == 0 {
			fmt.Println("[ERROR]: Division by zero")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f", a, b, a / b)
		}
	default:
		fmt.Println("Invalid operation")
	}
}