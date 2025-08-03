package main

import (
	"cmp"
	"fmt"
)

// Write the binary search algorithm. The function should take a sorted slice and a target element,
// and return the index of the element or -1 if it is not found.

func BinarySearch[T cmp.Ordered](arr []T, key T) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if cmp.Less(key, arr[mid]) {
			right = mid - 1
		} else if cmp.Less(arr[mid], key) {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := 7
	fmt.Println("Input:", arr)
	fmt.Println("Key:", key)

	idx := BinarySearch(arr, key)
	if idx != -1 {
		fmt.Println("Found key at index:", idx)
	} else {
		fmt.Println("Key not found")
	}
}
