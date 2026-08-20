package main

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Написать программу, которая считывает текст из файла input.txt, 
// подсчитывает количество слов в нем и записывает результат в файл output.txt.

func CountWords(nameFile string) int {
	contextBytes, _ := os.ReadFile(nameFile)
	context := string(contextBytes)

	words := strings.Fields(context)
	count := len(words)

	return count
}

func WriteContextToFile(context string) {
	os.WriteFile("output.txt", []byte(context), 0644)
}

func main() {
	count := CountWords("input.txt")
	str := strconv.Itoa(count)
	WriteContextToFile(str)
}