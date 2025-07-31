package main

import (
	"math"
	"testing"
)

func TestSetBit(t *testing.T) {
	type args struct {
		num   int64
		i     uint
		value bool
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Set 1th bit in '0' to 1",
			args: args{
				num:   0,
				i:     0,
				value: true,
			},
			want: 1,
		},
		{
			name: "Set 1th bit in '5' to 0",
			args: args{
				num:   5,
				i:     0,
				value: false,
			},
			want: 4,
		},
		{
			name: "Set 3rd bit in '0' to 1",
			args: args{
				num:   0,
				i:     2,
				value: true,
			},
			want: 4,
		},
		{
			name: "Set 64th bit in '0' to 1",
			args: args{
				num:   0,
				i:     63,
				value: true,
			},
			want: math.MinInt64,
		},
		{
			name: "Set 64th bit in '-1' to 0",
			args: args{
				num:   -1,
				i:     63,
				value: false,
			},
			want: math.MaxInt64,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetBit(tt.args.num, tt.args.i, tt.args.value); got != tt.want {
				t.Errorf("SetBit() = %v, want %v", got, tt.want)
			}
		})
	}
}
