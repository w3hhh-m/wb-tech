package main

import (
	"fmt"
)

// Remove the i-th element from a slice. Demonstrate proper removal without memory leaks.

func RemoveAt[T any](arr []T, idx int) []T {
	if idx < 0 || idx >= len(arr) {
		return arr
	}
	arr = append(arr[:idx], arr[idx+1:]...)
	clear(arr[len(arr) : len(arr)+1])
	return arr[:len(arr):len(arr)]
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Input:", arr)
	arr = RemoveAt(arr, 1)
	fmt.Println("After remove:", arr)
}
