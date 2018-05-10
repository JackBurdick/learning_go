package main

import "fmt"

func main() {
	arrSlice := []float64{98, 93, 77, 82, 83}

	total := 0.0
	for _, val := range arrSlice {
		total += val
	}
	fmt.Printf("Average = %v\n", total/float64(len(arrSlice)))
}
