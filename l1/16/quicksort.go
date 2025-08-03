package main

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	var left, right []int

	for i := 0; i < len(arr); i++ {
		if i == mid {
			continue
		}
		if arr[i] < arr[mid] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(append(QuickSort(left), arr[mid]), QuickSort(right)...)
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Println("Input:", arr)
	arr = QuickSort(arr)
	fmt.Println("Sorted:", arr)
}
