package main

import (
	"fmt"
	"math/rand"
)

func findMinMaxIndices(slice []int) (int, int) {
	minIndex, maxIndex := 0, 0

	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[minIndex] {
			minIndex = i
		}
		
		if slice[i] > slice[maxIndex] {
			maxIndex = i
		}
	}

	return minIndex, maxIndex
}

func main() {
	var length int

	fmt.Print("Enter length for create: ")
	fmt.Scan(&length)

	if length <= 0 {
		fmt.Println("Length must be greater than 0")
		return
	}

	slice := make([]int, 0, length)

	for i := 0; i < length; i++ {
		number := rand.Intn(50)
		slice = append(slice, number)
	}

	fmt.Println("Original slice:", slice)

	minIdx, maxIdx := findMinMaxIndices(slice)

	fmt.Printf("Min element: %d (index %d)\n", slice[minIdx], minIdx)
	fmt.Printf("Max element: %d (index %d)\n", slice[maxIdx], maxIdx)

	slice[minIdx], slice[maxIdx] = slice[maxIdx], slice[minIdx]

	fmt.Println("Modified slice:", slice)
}