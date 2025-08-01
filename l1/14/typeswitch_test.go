package main

import "testing"

func TestDetermineType(t *testing.T) {
	tests := []struct {
		name string
		v    any
		want string
	}{
		{
			name: "Integer type",
			v:    100,
			want: "int",
		},
		{
			name: "String type",
			v:    "text",
			want: "string",
		},
		{
			name: "Boolean type",
			v:    true,
			want: "bool",
		},
		{
			name: "Channel type",
			v:    make(chan int),
			want: "chan",
		},
		{
			name: "Unknown type - float64",
			v:    3.14,
			want: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetermineType(tt.v); got != tt.want {
				t.Errorf("DetermineType() = %v, want %v", got, tt.want)
			}
		})
	}
}
