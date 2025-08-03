package main

import (
	"fmt"
	"strings"
)

// Write a program that reverses the words order in input string.

func ReverseWords(s string) string {
	arr := strings.Fields(s)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return strings.Join(arr, " ")
}

func main() {
	s := "snow dog sun"
	fmt.Println("Input:", s)
	fmt.Println("Reversed:", ReverseWords(s))
}
