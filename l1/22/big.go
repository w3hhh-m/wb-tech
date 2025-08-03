package main

import (
	"fmt"
	"math/big"
)

// Write a program that multiplies, divides, adds, and subtracts
// two integer variables a and b, whose values are greater than 2^20

func main() {
	a := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(60), nil) // a = 2^60
	b := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(30), nil) // b = 2^30

	sum := big.NewInt(0).Add(a, b)
	diff := big.NewInt(0).Sub(a, b)
	prod := big.NewInt(0).Mul(a, b)

	fmt.Println("a =", a.String())
	fmt.Println("b =", b.String())
	fmt.Println("sum =", sum.String())
	fmt.Println("diff =", diff.String())
	fmt.Println("prod =", prod.String())

	// If b == 0, a division-by-zero run-time panic occurs.
	if b.Sign() != 0 {
		// Quo implements truncated division (like Go)
		// Example: -11 / 4 = -2
		// -11 = 4*(-2) + (-3)
		// The sign of the remainder is the same as the sign of the dividend
		quo := big.NewInt(0).Quo(a, b)

		// Div implements Euclidean division (unlike Go)
		// Example: -11 / 4 = -3
		// -11 = 4*(-3) + 1
		// The remainder must be positive
		div := big.NewInt(0).Div(a, b)

		fmt.Println("quo =", quo.String())
		fmt.Println("div =", div.String())
	}
}
