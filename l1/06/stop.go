package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Implement all possible ways to stop a goroutine execution.
// Classic solutions: exit by condition, via notification channel, via context, terminating with runtime.Goexit(), etc.

func main() {
	var wg sync.WaitGroup

	// === via notification channel ===
	doneCh := make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		workerChannel(doneCh)
	}()
	// work imitation
	time.Sleep(1 * time.Second)
	// exiting goroutine
	doneCh <- true
	// will wait until exit
	wg.Wait()

	// === via context ===
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		workerContext(ctx)
	}()
	// work imitation
	time.Sleep(1 * time.Second)
	// exiting goroutine
	cancel()
	// will wait until exit
	wg.Wait()

	// === runtime.Goexit() ===
	wg.Add(1)
	go func() {
		defer wg.Done()
		workerRuntimeGoexit()
	}()
	// will wait until exit
	wg.Wait()

	// === channel close ===
	stopChan := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		workerChannelClose(stopChan)
	}()
	// work imitation
	time.Sleep(1 * time.Second)
	// exiting goroutine
	close(stopChan)
	// will wait until exit
	wg.Wait()

	// === timeout ===
	wg.Add(1)
	go func() {
		defer wg.Done()
		workerTimeout(1 * time.Second)
	}()
	// will wait until exit
	wg.Wait()
}

func workerChannel(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("[doneCh] Stopped via done channel signal")
			return
		default:
			fmt.Println("[doneCh] Working...")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func workerContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[context] Stopped via context cancellation")
			return
		default:
			fmt.Println("[context] Working...")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func workerRuntimeGoexit() {
	defer fmt.Println("[Goexit] Stopped via runtime.Goexit()")

	fmt.Println("[Goexit] Working...")
	time.Sleep(200 * time.Millisecond)
	runtime.Goexit()
}

func workerChannelClose(stopChan chan struct{}) {
	for {
		select {
		case _, ok := <-stopChan:
			if !ok {
				fmt.Println("[channel close] Stopped via channel close")
				return
			}
		default:
			fmt.Println("[channel close] Working...")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func workerTimeout(timeout time.Duration) {
	deadline := time.After(timeout)
	for {
		select {
		case <-deadline:
			fmt.Println("[timeout] Stopped via timeout")
			return
		default:
			fmt.Println("[timeout] Working...")
			time.Sleep(200 * time.Millisecond)
		}
	}
}
