package main

import "fmt"

// Write a program that reverses the input string.

func ReverseString(s string) string {
	arr := []rune(s)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr)
}

func main() {
	s := "главрыба"
	fmt.Println("Input:", s)
	fmt.Println("Reversed:", ReverseString(s))
}
