package main

import (
	"fmt"
	"time"
)

// Create your own sleep(duration) function that works similarly to the built-in time.Sleep function.

func Sleep(duration time.Duration) {
	<-time.After(duration)

	// same:
	// timer := time.NewTimer(duration)
	// <-timer.C
}

func BusySleep(duration time.Duration) {
	end := time.Now().Add(duration)
	for time.Now().Before(end) {
		// cpu spin
	}
}

func main() {
	start := time.Now()
	Sleep(time.Second)
	fmt.Println("Sleep took", time.Since(start))

	start = time.Now()
	BusySleep(time.Second)
	fmt.Println("BusySleep took", time.Since(start))
}
