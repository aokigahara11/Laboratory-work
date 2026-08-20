package main

import (
	"fmt"
)

// Найти сумму всех четных и произведение всех нечетных чисел в диапазоне от $1$ до $N$.

func main() {
	var n int
	var sumEven int
	var productOdd int = 1

	fmt.Print("Enter a positive integer N: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i % 2 == 0 {
			sumEven += i
		} else {
			productOdd *= i
		}
	}
	
	fmt.Printf("Sum of even numbers: %d\n", sumEven)
	fmt.Printf("Product of odd numbers: %d\n", productOdd)
}