package main

import (
	"context"
	"fmt"
	"time"
)

// Create a program that sends values into a channel one by one, and on the other side of the channel reads these values.
// The program should stop after N seconds.

func main() {
	data := "some data"
	timeout := 2 * time.Second

	dataCh := make(chan string, 1)
	defer close(dataCh)

	// two ways with channel and context
	timeoutCh := time.After(timeout)
	timeoutCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		select {
		case <-timeoutCh:
			// there is no guarantee that this case will run exactly after the timeout
			// because select chooses cases randomly when several are ready
			fmt.Println("timed out channel")
			return
		case <-timeoutCtx.Done():
			// there is no guarantee that this case will run exactly after the timeout
			// because select chooses cases randomly when several are ready
			fmt.Println("timed out context")
			return
		case dataCh <- data:
			fmt.Println("sent data:", data)
		case val := <-dataCh:
			fmt.Println("received data:", val)
		}
	}
}
