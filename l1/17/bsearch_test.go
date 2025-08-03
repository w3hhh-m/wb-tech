package main

import (
	"cmp"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name string
		arr  []T
		key  T
		want int
	}

	intTests := []testCase[int]{
		{
			name: "Found",
			arr:  []int{1, 2, 3, 4, 5},
			key:  3,
			want: 2,
		},
		{
			name: "Not found",
			arr:  []int{1, 2, 3, 4, 5},
			key:  6,
			want: -1,
		},
		{
			name: "Empty array",
			arr:  []int{},
			key:  0,
			want: -1,
		},
	}

	stringTests := []testCase[string]{
		{
			name: "Found",
			arr:  []string{"apple", "banana", "cherry"},
			key:  "apple",
			want: 0,
		},
		{
			name: "Not found",
			arr:  []string{"apple", "banana", "cherry"},
			key:  "melon",
			want: -1,
		},
		{
			name: "Empty array",
			arr:  []string{},
			key:  "",
			want: -1,
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.arr, tt.key); got != tt.want {
				t.Errorf("[INTEGER] BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.arr, tt.key); got != tt.want {
				t.Errorf("[STRING] BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
