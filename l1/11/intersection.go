package main

import "fmt"

// Implement the intersection of two unordered sets (e.g., two slices)
// i.e., output the elements that are present in both the first and the second set.
// Example:
// A = {1,2,3}
// B = {2,3,4}
// Intersection = {2,3}

func FindIntersection(a, b []int) []int {
	aMap := make(map[int]bool)
	for _, val := range a {
		aMap[val] = true
	}

	intersection := []int{}

	for _, val := range b {
		if aMap[val] {
			intersection = append(intersection, val)
		}
	}

	return intersection
}

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	intersection := FindIntersection(a, b)

	fmt.Printf("A = %v\n", a)
	fmt.Printf("B = %v\n", b)
	fmt.Printf("Intersection = %v\n", intersection)
}
