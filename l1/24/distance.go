package main

import (
	"fmt"
	"math"
)

// Write a program to calculate the distance between two points on a plane.
// The points are represented as a Point structure with encapsulated (private) fields
// x, y (of type float64) and a constructor.
// The distance is calculated using the formula between the coordinates of two points.

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(1, 0)
	d := p1.Distance(p2)

	fmt.Println("P1:", p1)
	fmt.Println("P2:", p2)
	fmt.Println("Distance:", d)
}
