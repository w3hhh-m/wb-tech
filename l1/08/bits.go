package main

import "fmt"

// Given a variable of type int64, write a program that sets the i-th bit of this number to either 1 or 0

func SetBit(num int64, i uint, value bool) int64 {
	if value {
		return num | (1 << i)
	} else {
		return num & ^(1 << i)
	}
}

func main() {
	var num int64 = 5 // Binary: 0101
	fmt.Printf("Original: %d (%b)\n", num, num)
	// Set 1st bit to 0
	result := SetBit(num, 0, false)
	fmt.Printf("Set 1st bit to 0: %d (%b)\n", result, result)
}
