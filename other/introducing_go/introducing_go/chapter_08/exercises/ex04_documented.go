package main

import "fmt"

// returns max in a series
func Max(xs []float64) float64 {
	var maxVal float64
	for i, x := range xs {
		if i == 0 || x > maxVal {
			maxVal = x
		}
	}
	return maxVal
}

// returns min in a series
func Min(xs []float64) float64 {
	var minVal float64
	for i, x := range xs {
		if i == 0 || x < minVal {
			minVal = x
		}
	}
	return minVal
}

func main() {
	exSlice := []float64{98, 93, 77, 82, 83}
	fmt.Println(Max(exSlice))
	fmt.Println(Min(exSlice))
}
