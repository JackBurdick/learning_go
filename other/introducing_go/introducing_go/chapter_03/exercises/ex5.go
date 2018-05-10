package main

import "fmt"

func main() {
	fmt.Print("Enter F* to convert to C*: ")
	var input float64
	var output float64

	fmt.Scanf("%f", &input)
	// convert to C
	output = (input - 32) * 5 / 9

	outputString := fmt.Sprintf("%v F* = %v C*", input, output)
	fmt.Println(outputString)
}
