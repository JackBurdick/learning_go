package main

import "fmt"

func returnTwo() (int, int) {
	return 4, 9
}

func main() {
	x, y := returnTwo()
	fmt.Printf("x = %v, y = %v\n", x, y)
}
