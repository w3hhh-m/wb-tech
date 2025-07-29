package main

import (
	"fmt"
	"sync"
)

// Write a program that concurrently calculates the squares of numbers
// from the array [2, 4, 6, 8, 10] and prints the results to stdout.

func main() {
	nums := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, num := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(num * num)
		}()
	}

	wg.Wait()
}
