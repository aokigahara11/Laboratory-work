package main

import (
	"fmt"
)

// Вычислить факториал числа $N$ с помощью цикла.

func main() {
	var n int
	fmt.Print("Enter a positive integer N: ")
	fmt.Scan(&n)

	var factorial int64 = 1
	for i := 1; i <= n; i++ {
		factorial *= int64(i)
	}

	fmt.Printf("Factorial of %d is %d\n", n, factorial)
}