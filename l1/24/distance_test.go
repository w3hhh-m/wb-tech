package main

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		name string
		p1   Point
		p2   Point
		want float64
	}{
		{
			name: "Same points",
			p1:   NewPoint(0, 0),
			p2:   NewPoint(0, 0),
			want: 0,
		},
		{
			name: "Horizontal distance",
			p1:   NewPoint(0, 0),
			p2:   NewPoint(3, 0),
			want: 3,
		},
		{
			name: "Vertical distance",
			p1:   NewPoint(0, 0),
			p2:   NewPoint(0, 4),
			want: 4,
		},
		{
			name: "Diagonal distance",
			p1:   NewPoint(1, 1),
			p2:   NewPoint(4, 5),
			want: 5,
		},
		{
			name: "Negative coordinates",
			p1:   NewPoint(-1, -1),
			p2:   NewPoint(-4, -5),
			want: 5,
		},
		{
			name: "Mixed coordinates",
			p1:   NewPoint(-2, 3),
			p2:   NewPoint(1, -1),
			want: 5,
		},
		{
			name: "Float coordinates",
			p1:   NewPoint(1.5, 2.5),
			p2:   NewPoint(4.5, 6.5),
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p1.Distance(tt.p2); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
