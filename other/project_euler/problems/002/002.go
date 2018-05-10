package main

import "fmt"

func main() {
	// specification; value must be under 4M
	MAXVALUE := 4000000
	var fibSeq []int
	var evenSum int

	// build fib sequence
	// initialize with first two values
	fibSeq = append(fibSeq, 1, 2)
	evenSum += 2
	nextVal := sumLastTwoValues(fibSeq)
	// appended new values provided new value is under spec
	for nextVal < MAXVALUE {
		fibSeq = append(fibSeq, nextVal)
		nextVal = sumLastTwoValues(fibSeq)
		if nextVal % 2 == 0 {
			evenSum += nextVal
		}
	}
	fmt.Println(evenSum)

	// DEBUG: confirm fib sequence
	printSlice(fibSeq)

}

// add the last two items in the slice
func sumLastTwoValues(fibSeq []int) int {
	fsLastIndex := len(fibSeq) - 1
	valA := fibSeq[fsLastIndex -1]
	valB := fibSeq[fsLastIndex]
	valNew := valA + valB
	return valNew
}

// display slice information
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}