package main

import (
	"fmt"
)

// Посчитать количество цифр и сумму цифр в заданном натуральном числе $N$.

func main() {
	var n int

	fmt.Print("Enter a positive integer N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("The number must be positive.")
		return
	}

	temp := n
	var count int
	var sum int

	for temp > 0 {
		digit := temp % 10
		count++
		sum += digit
		temp /= 10
	}

	fmt.Printf("Number of digits: %d\n", count)
	fmt.Printf("Sum of digits: %d\n", sum)
}