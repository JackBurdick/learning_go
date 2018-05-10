package main

import "fmt"

// recursive fibonacci calls
func fibonacci(val int) int {
	if val == 1 || val == 0 {
		return val
	} else {
		return (fibonacci(val-1) + fibonacci(val-2))
	}
}

func main() {
	for i := 0; i < 15; i++ {
		fibVal := fibonacci(i)
		fmt.Printf("fibonacci(%v)\t = %v\n", i, fibVal)
	}
}
