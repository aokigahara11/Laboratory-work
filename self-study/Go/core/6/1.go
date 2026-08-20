package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {
	str = strings.ToLower(str)

	runes := []rune(str)

	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("Dota"))
	fmt.Println(isPalindrome("Топот"))
	fmt.Println(isPalindrome("hello"))
}