package main

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Default example",
			s:    "snow dog sun",
			want: "sun dog snow",
		},
		{
			name: "Empty string",
			s:    "",
			want: "",
		},
		{
			name: "Single word",
			s:    "word",
			want: "word",
		},
		{
			name: "With extra spaces",
			s:    "   one   two three  ",
			want: "three two one",
		},
		{
			name: "Emoji",
			s:    "🙏 ✅ 🎉",
			want: "🎉 ✅ 🙏",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseWords(tt.s); got != tt.want {
				t.Errorf("ReverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
