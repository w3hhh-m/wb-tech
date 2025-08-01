package main

import (
	"reflect"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{
			name: "Basic example",
			a:    []int{1, 2, 3},
			b:    []int{2, 3, 4},
			want: []int{2, 3},
		},
		{
			name: "No intersection",
			a:    []int{1, 2, 3},
			b:    []int{4, 5, 6},
			want: []int{},
		},
		{
			name: "Complete intersection",
			a:    []int{1, 2, 3},
			b:    []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "Empty slices",
			a:    []int{},
			b:    []int{},
			want: []int{},
		},
		{
			name: "One empty slice",
			a:    []int{1, 2, 3},
			b:    []int{},
			want: []int{},
		},
		{
			name: "Duplicate elements",
			a:    []int{1, 2, 2, 3, 3},
			b:    []int{2, 2, 3, 4},
			want: []int{2, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindIntersection(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
