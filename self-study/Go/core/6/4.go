package main

import (
	"fmt"
	"os"
	"strings"
)

// Написать функцию, которая принимает путь к файлу и 
// заменяет в нем все вхождения заданного слова $A$ на слово $B$, перезаписывая файл.

func ReplaceWord(newWord string, oldWord string, fileName string) {
	contextBytes, _ := os.ReadFile(fileName)
	content := string(contextBytes)

	updatedText := strings.ReplaceAll(content, oldWord, newWord)
	os.WriteFile(fileName, []byte(updatedText), 0644)
}

func main() {
	var oldWord, newWord, fileName string
	
	fmt.Print("Enter words for replace in file (Old/New): ")
	fmt.Scan(&oldWord, &newWord)
	
	fmt.Println("Enter name file: ")
	fmt.Scan(&fileName)

	ReplaceWord(newWord, oldWord, fileName)
}