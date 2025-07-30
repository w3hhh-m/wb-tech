package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// Continuously write data into a channel from the main goroutine.
// NOTE: I removed infinite writing and graceful shutdown because it is in the next task
// Create a set of N worker goroutines that read data from this channel.
// Each worker should print the data it reads to stdout.
// The program should accept the number of workers as a parameter and start that many worker goroutines.

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run workers.go <number_of_workers>")
		os.Exit(1)
	}

	workers, err := strconv.Atoi(os.Args[1])
	if err != nil || workers < 1 {
		fmt.Println("Invalid number of workers")
		os.Exit(1)
	}

	ch := make(chan any)
	var wg sync.WaitGroup

	// writing goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()

		// closing channel when exiting
		defer close(ch)
		for i := 0; i < 100; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// workers
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// worker will stop when chan is closed
			for data := range ch {
				fmt.Printf("Worker %d: %v\n", i, data)
			}
		}()
	}

	// waiting for all workers to exit
	wg.Wait()
}
