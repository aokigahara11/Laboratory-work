package main

import (
	"fmt"
	"strings"
)

// Создать карту (map), где ключом является слово, а значением — количество его повторений в некотором тексте.

func wordCount(text string) map[string]int {
	wordMap := make(map[string]int)
	var wordCount int

	words := strings.Fields(text)

	for _, word := range words {
		cleanedWord := strings.ToLower(word)
		wordMap[cleanedWord]++
	}
	
	return wordMap
}

func main() {
	fmt.Print("Enter text to count word occurrences: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()

		result := wordCount(text)
		fmt.Println("Word counts:", result)
	}
}