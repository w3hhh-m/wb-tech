package main

import "fmt"

// Write a program that uses two channels:
// the first channel receives numbers x from an array,
// and the second channel receives the result of the operation x * 2.
// The data from the second channel should be printed to stdout.
// Organize a two-stage pipeline with goroutines:
// one for generating numbers and another for processing them.
// Ensure that reading from the second channel completes correctly.

func main() {
	nums := []int{1, 2, 3, 4, 5}

	genChan := make(chan int)
	doubleChan := make(chan int)

	// generate numbers
	go func() {
		defer close(genChan)
		for _, x := range nums {
			genChan <- x
		}
	}()

	// double numbers
	go func() {
		defer close(doubleChan)
		for x := range genChan {
			doubleChan <- x * 2
		}
	}()

	// not using waitgroups because main goroutine
	// will anyway block until 2nd goroutine exits
	// and 2nd goroutine will block until 1st goroutine exits

	// read doubled
	for doubled := range doubleChan {
		fmt.Println(doubled)
	}
}
