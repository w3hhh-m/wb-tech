package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Default example",
			arr:  []int{3, 1, 4, 1, 5, 9},
			want: []int{1, 1, 3, 4, 5, 9},
		},
		{
			name: "Empty array",
			arr:  []int{},
			want: []int{},
		},
		{
			name: "Single element",
			arr:  []int{1},
			want: []int{1},
		},
		{
			name: "Reverse order",
			arr:  []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Already sorted",
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
