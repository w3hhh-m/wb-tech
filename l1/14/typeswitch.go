package main

import "fmt"

// Write a program that can determine the type of variable
// passed to it at runtime (an interface{} is given as input).
// The types to be recognized are: int, string, bool, and chan (channel)

func DetermineType(v any) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int, chan string, chan bool:
		return "chan"
	default:
		return "unknown"
	}
}

func main() {
	intVar := 100
	strVar := "text"
	boolVar := true
	chanVar := make(chan int)

	fmt.Printf("Type of %v is: %s\n", intVar, DetermineType(intVar))
	fmt.Printf("Type of %v is: %s\n", strVar, DetermineType(strVar))
	fmt.Printf("Type of %v is: %s\n", boolVar, DetermineType(boolVar))
	fmt.Printf("Type of %v is: %s\n", chanVar, DetermineType(chanVar))

	floatVar := 3.14

	fmt.Printf("Type of %v is: %s\n", floatVar, DetermineType(floatVar))
}
