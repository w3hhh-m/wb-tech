package main

import (
	"fmt"
	"os"
	"wb-tech-l0/internal/app"
)

// main is the entry point of the application
func main() {
	if err := app.RunApp(); err != nil {
		fmt.Printf("Application error: %s\n", err)
		os.Exit(1)
	}
}
