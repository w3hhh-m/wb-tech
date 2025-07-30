package main

import (
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

	timeoutCh := time.After(timeout)

	for {
		select {
		case <-timeoutCh:
			// there is no guarantee that this case will run exactly after the timeout
			// because select chooses cases randomly when several are ready
			fmt.Println("timed out")
			return
		case dataCh <- data:
			fmt.Println("sent data:", data)
		case val := <-dataCh:
			fmt.Println("received data:", val)
		}
	}
}
