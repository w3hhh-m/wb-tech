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

// Task 3 with additional condition:
// The program should terminate correctly when pressed ctrl+C.

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
				fmt.Println("Writer: context canceled. Closing channel")
				return
			case ch <- i:
				// sleep with also waiting for exiting
				select {
				case <-ctx.Done():
					fmt.Println("Writer: context canceled while sleeping. Closing channel")
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
			fmt.Printf("Worker %d: channel closed. Stop reading\n", i)
		}()
	}

	// waiting for all workers to exit
	wg.Wait()
	fmt.Println("Shutdown complete")
}
