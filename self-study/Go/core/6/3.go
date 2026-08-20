package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Задана строка, содержащая числа, разделенные запятыми (например, "10,20,30,40"). 
// Распарсить строку с помощью strings.Split и strconv.Atoi, найдя сумму всех чисел.

func GetSumDigits(numbersStr string) int {
	numbers := strings.Split(numbersStr, ",")
	sum := 0

	for _, word := range numbers {
		digit, _ := strconv.Atoi(word)  
		sum += digit
	}
	return sum
}

func main() {
	numbers := "10,20,30,40"
	answer := GetSumDigits(numbers)
	fmt.Println("Sum numbers:", answer)
}