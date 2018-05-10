package main

import "fmt"

func findMax(vals ...int) int {
	var maxVal int
	for i, val := range vals {
		// if first round, keep value, check against max
		if i == 0 || val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func main() {
	fmt.Printf("Max val = %v\n", findMax(2, 4, 5, 6))
}
