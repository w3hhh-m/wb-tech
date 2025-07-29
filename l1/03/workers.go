package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// Continuously write data into a channel from the main goroutine.
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

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	ch := make(chan any)
	// writing goroutine
	go func() {
		// closing channel when exiting
		defer close(ch)
		i := 0
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
				// sleep with also waiting for exiting
				select {
				case <-ctx.Done():
					return
				case <-time.After(100 * time.Millisecond):
					i++
				}
			}
		}
	}()

	var wg sync.WaitGroup
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
