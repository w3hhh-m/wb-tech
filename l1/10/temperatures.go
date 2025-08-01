package main

import "fmt"

// Given a sequence of temperature fluctuations: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5,
// group these values into buckets with a step of 10 degrees.
// Example:
// -20: {-25.4, -27.0, -21.0}
// 10: {13.0, 19.0, 15.5}
// 20: {24.5}
// 30: {32.5}

// It seems to me that the task is not entirely correct,
// because there will either be two groups: -0 and 0,
// or in group 0 there will be values from -10 to 10.
// But who am I to change the task :)

func GroupTemperatures(temps []float64) map[int][]float64 {
	groups := make(map[int][]float64)
	for _, temp := range temps {
		group := int(temp/10) * 10
		groups[group] = append(groups[group], temp)
	}
	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := GroupTemperatures(temperatures)

	for k, v := range groups {
		fmt.Printf("%d: %v\n", k, v)
	}
}
