package main

import "fmt"

// Swap the values of two integer variables
// without using a temporary variable.

func SwapUsingMath(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func SwapUsingXOR(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func main() {
	a, b := 10, 20
	fmt.Printf("Before math swap: a = %d, b = %d\n", a, b)
	SwapUsingMath(&a, &b)
	fmt.Printf("After math swap: a = %d, b = %d\n\n", a, b)

	c, d := 30, 40
	fmt.Printf("Before XOR swap: c = %d, d = %d\n", c, d)
	SwapUsingXOR(&c, &d)
	fmt.Printf("After XOR swap: c = %d, d = %d\n", c, d)
}
