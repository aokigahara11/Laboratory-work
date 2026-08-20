package main

import (
	"fmt"
)

// Написать функцию, которая принимает слайс целых чисел и возвращает новый слайс, 
// содержащий  только уникальные элементы (используйте map для фильтрации).

func uniqueElements(slice []int) []int {
	uniqueMap := make(map[int]bool)
	uniqueSlice := []int{}
	
	for _, num := range slice {
		if _, exists := uniqueMap[num]; !exists {
			uniqueMap[num] = true
			uniqueSlice = append(uniqueSlice, num)
		}
	}
	return uniqueSlice
}