package main

import "testing"

func TestSwapUsingMath(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
	}{
		{
			name: "Positive numbers",
			a:    5,
			b:    10,
		},
		{
			name: "Zero and positive",
			a:    0,
			b:    7,
		},
		{
			name: "Negative numbers",
			a:    -3,
			b:    -8,
		},
		{
			name: "Positive and negative",
			a:    4,
			b:    -9,
		},
		{
			name: "Same numbers",
			a:    6,
			b:    6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b := tt.a, tt.b
			originalA, originalB := a, b

			SwapUsingMath(&a, &b)

			if a != originalB || b != originalA {
				t.Errorf("SwapUsingMath() = %v, %v; want %v, %v", a, b, originalB, originalA)
			}
		})
	}
}

func TestSwapUsingXOR(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
	}{
		{
			name: "Positive numbers",
			a:    5,
			b:    10,
		},
		{
			name: "Zero and positive",
			a:    0,
			b:    7,
		},
		{
			name: "Negative numbers",
			a:    -3,
			b:    -8,
		},
		{
			name: "Positive and negative",
			a:    4,
			b:    -9,
		},
		{
			name: "Same numbers",
			a:    6,
			b:    6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b := tt.a, tt.b
			originalA, originalB := a, b

			SwapUsingXOR(&a, &b)

			if a != originalB || b != originalA {
				t.Errorf("SwapUsingXOR() = %v, %v; want %v, %v", a, b, originalB, originalA)
			}
		})
	}
}
