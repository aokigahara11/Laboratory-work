package main 

import (
	"fmt"
)

// Создать функцию, определяющую, является ли переданное число простым.

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&n)

	if IsPrime(n) {
		fmt.Printf("%d is a prime number.\n", n)
	} else {
		fmt.Printf("%d is not a prime number.\n", n)
	}
}