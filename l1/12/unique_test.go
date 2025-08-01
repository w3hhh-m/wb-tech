package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGetUniqueStrings(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "Empty input",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "No duplicates",
			input: []string{"cat", "dog", "tree"},
			want:  []string{"cat", "dog", "tree"},
		},
		{
			name:  "With duplicates",
			input: []string{"cat", "cat", "dog", "cat", "tree"},
			want:  []string{"cat", "dog", "tree"},
		},
		{
			name:  "All duplicates",
			input: []string{"cat", "cat", "cat"},
			want:  []string{"cat"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetUniqueStrings(tt.input)

			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUniqueStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
