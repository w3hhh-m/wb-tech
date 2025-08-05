package main

import "testing"

func TestAreSymbolsUnique(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Empty string",
			s:    "",
			want: true,
		},
		{
			name: "Single symbol",
			s:    "a",
			want: true,
		},
		{
			name: "All unique",
			s:    "abc",
			want: true,
		},
		{
			name: "Duplicate symbol",
			s:    "abca",
			want: false,
		},
		{
			name: "Emoji",
			s:    "🙏🙏",
			want: false,
		},
		{
			name: "Numbers duplicate",
			s:    "1123",
			want: false,
		},
		{
			name: "Spaces duplicate",
			s:    "1 2 3",
			want: false,
		},
		{
			name: "Case duplicates",
			s:    "abcA",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreSymbolsUnique(tt.s); got != tt.want {
				t.Errorf("AreSymbolsUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
