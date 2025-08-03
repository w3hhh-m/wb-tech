package main

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Default example",
			s:    "главрыба",
			want: "абырвалг",
		},
		{
			name: "Empty string",
			s:    "",
			want: "",
		},
		{
			name: "Single letter",
			s:    "a",
			want: "a",
		},
		{
			name: "Emoji",
			s:    "🙏✅🎉",
			want: "🎉✅🙏",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.s); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
