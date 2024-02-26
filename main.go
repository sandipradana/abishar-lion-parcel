package main

import "fmt"

func Palindrome(word string) bool {
	reversed := ""

	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}

	return word == reversed
}

func main() {
	fmt.Println("Test 1")
	fmt.Printf("radar is palindrome? %t\n", Palindrome("radar"))
	fmt.Printf("hello is palindrome? %t\n", Palindrome("hello"))
}
