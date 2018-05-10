package main

import "fmt"

func main() {
	arr := []int{
		49, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	// arr := []int{}

	// ensure arr exists
	if arr != nil {
		// still would need to check type
		var curMin int
		curMin = arr[0]
		// iterate though, save curMin
		for _, num := range arr {
			if num < curMin {
				curMin = num
			}
		}
		fmt.Printf("Min: %v\n", curMin)
	} else {
		fmt.Println("arr == nil")
	}
}
