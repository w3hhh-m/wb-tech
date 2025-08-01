package main

import "fmt"

// Create a set (collection of unique elements) from a given
// sequence of strings: ("cat", "cat", "dog", "cat", "tree").
// The expected result is a set of unique words: {"cat", "dog", "tree"}.

func GetUniqueStrings(strings []string) []string {
	m := make(map[string]struct{})

	for _, s := range strings {
		m[s] = struct{}{}
	}

	unique := make([]string, 0, len(m))
	for s := range m {
		unique = append(unique, s)
	}

	return unique
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	unique := GetUniqueStrings(strings)
	fmt.Println("Unique strings:", unique)
}
