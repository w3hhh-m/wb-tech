package main

import (
	"reflect"
	"testing"
)

func TestRemoveAt(t *testing.T) {
	type testCase[T any] struct {
		name string
		arr  []T
		idx  int
		want []T
	}

	testsInt := []testCase[int]{
		{
			name: "Default example",
			arr:  []int{1, 2, 3, 4, 5},
			idx:  1,
			want: []int{1, 3, 4, 5},
		},
		{
			name: "Empty example",
			arr:  []int{},
			idx:  0,
			want: []int{},
		},
		{
			name: "Single item example",
			arr:  []int{1},
			idx:  0,
			want: []int{},
		},
		{
			name: "Out of range example",
			arr:  []int{1, 2, 3, 4, 5},
			idx:  8,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Negative index example",
			arr:  []int{1, 2, 3, 4, 5},
			idx:  -2,
			want: []int{1, 2, 3, 4, 5},
		},
	}

	testsString := []testCase[string]{
		{
			name: "Default example",
			arr:  []string{"1", "2", "3", "4", "5"},
			idx:  1,
			want: []string{"1", "3", "4", "5"},
		},
		{
			name: "Empty example",
			arr:  []string{},
			idx:  0,
			want: []string{},
		},
		{
			name: "Single item example",
			arr:  []string{"1"},
			idx:  0,
			want: []string{},
		},
		{
			name: "Out of range example",
			arr:  []string{"1", "2", "3", "4", "5"},
			idx:  8,
			want: []string{"1", "2", "3", "4", "5"},
		},
		{
			name: "Negative index example",
			arr:  []string{"1", "2", "3", "4", "5"},
			idx:  -2,
			want: []string{"1", "2", "3", "4", "5"},
		},
	}

	for _, tt := range testsInt {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveAt(tt.arr, tt.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[INTEGER] RemoveAt() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range testsString {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveAt(tt.arr, tt.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[STRING] RemoveAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
