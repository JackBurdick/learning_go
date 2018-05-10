package main

import "fmt"

func factorial(x uint) uint {
	// base case
	if x == 0 {
		return 1
	}
	// call itself w/lesser value(closer to base case)
	return x * factorial(x-1)
}

func main() {
	fmt.Printf("factorial(2) = %v\n", factorial(2))
	fmt.Printf("factorial(7) = %v\n", factorial(7))
}
