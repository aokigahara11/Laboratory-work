package main

import (
	"fmt"
)

// Написать вариативную функцию minMax(nums ...int) (int, int), 
// возвращающую наименьшее и наибольшее число из переданных аргументов.

func minMax(nums ...int) (int, int) {
	min := nums[0]
	max := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

func main() {
	var count int
	fmt.Print("Enter count digit: ")
	fmt.Scan(&count)

	// Собираем все числа в слайс
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&numbers[i])
	}
	
	min, max := minMax(numbers...)
	fmt.Printf("Min: %d, Max: %d\n", min, max)
}