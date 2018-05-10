package main

import "fmt"

func add(args ...int) int {
	sum := 0
	for _, val := range args {
		sum += val
	}
	return sum
}

func main() {
	fmt.Println(add(1, 3, 5, 67))
}
