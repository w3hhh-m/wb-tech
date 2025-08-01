package main

import (
	"reflect"
	"testing"
)

func TestGroupTemperatures(t *testing.T) {
	tests := []struct {
		name  string
		temps []float64
		want  map[int][]float64
	}{
		{
			name:  "Empty input",
			temps: []float64{},
			want:  map[int][]float64{},
		},
		{
			name:  "Single temperature",
			temps: []float64{23.5},
			want: map[int][]float64{
				20: {23.5},
			},
		},
		{
			name:  "Multiple temperatures in one group",
			temps: []float64{21.5, 25.3, 27.8},
			want: map[int][]float64{
				20: {21.5, 25.3, 27.8},
			},
		},
		{
			name:  "Multiple temperatures in different groups",
			temps: []float64{-12.4, 3.6, 15.2, 24.8, 32.1},
			want: map[int][]float64{
				-10: {-12.4},
				0:   {3.6},
				10:  {15.2},
				20:  {24.8},
				30:  {32.1},
			},
		},
		{
			name:  "Corner cases",
			temps: []float64{-20.0, -10.0, -9.9, 0.0, 9.9, 10.0, 20.0},
			want: map[int][]float64{
				-20: {-20.0},
				-10: {-10.0},
				0:   {-9.9, 0.0, 9.9},
				10:  {10.0},
				20:  {20.0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupTemperatures(tt.temps)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}
