package main

import "fmt"

// functions form a call stack

func average(arrSlice []float64) float64 {
	total := 0.0
	for _, value := range arrSlice {
		total += value
	}
	average := total / float64(len(arrSlice))
	return average
}

func main() {
	exSlice := []float64{98, 93, 77, 82, 83}

	fmt.Printf("Average = %v\n", average(exSlice))
}
