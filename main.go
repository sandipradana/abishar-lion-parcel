package main

import "fmt"

func Palindrome(word string) bool {
	reversed := ""

	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}

	return word == reversed
}

func Max(numbers []int) int {
	max := numbers[0]

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}

func main() {
	fmt.Println("Test 1")
	fmt.Printf("radar is palindrome? %t\n", Palindrome("radar"))
	fmt.Printf("hello is palindrome? %t\n", Palindrome("hello"))

	fmt.Println("Test 2")
	fmt.Printf("max number from [3, 5, 1, 9, 2]  %d\n", Max([]int{3, 5, 1, 9, 2}))
	fmt.Printf("max number from [1, 2, 3, 4, 5]  %d\n", Max([]int{1, 2, 3, 4, 5}))
	fmt.Printf("max number from [5, 4, 3, 2, 1]  %d\n", Max([]int{5, 4, 3, 2, 1}))
}
