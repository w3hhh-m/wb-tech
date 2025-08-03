package main

import (
	"fmt"
	"io"
)

// Implement the Adapter design pattern using any example.

// Printer is a custom type with its own Print method
type Printer struct{}

func (p *Printer) Print(a ...any) {
	fmt.Print(a...)
}

// WriterAdapter adapts Printer to io.Writer interface
type WriterAdapter struct {
	printer *Printer
}

func (wa *WriterAdapter) Write(p []byte) (n int, err error) {
	// if we comment line below, nothing will be printed in main()
	wa.printer.Print(string(p))
	return len(p), nil
}

func main() {
	printer := &Printer{}
	adapter := &WriterAdapter{printer}

	var w io.Writer = adapter
	// fmt.Fprint writes to io.Writer, and through the adapter it prints using Printer
	fmt.Fprint(w, "Hello World!\n")
}
