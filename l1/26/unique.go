package main

import (
	"fmt"
	"strings"
)

// Write a program that checks if all characters in a string are unique.
// The check should be case-insensitive

func AreSymbolsUnique(s string) bool {
	m := make(map[rune]struct{})
	s = strings.ToLower(s)

	for _, char := range s {
		if _, ok := m[char]; ok {
			return false
		}
		m[char] = struct{}{}
	}
	return true
}

func main() {
	s := "abcd"
	fmt.Println("Input:", s)
	fmt.Println("Unique:", AreSymbolsUnique(s))

	s = "abcdA"
	fmt.Println("Input:", s)
	fmt.Println("Unique:", AreSymbolsUnique(s))
}
