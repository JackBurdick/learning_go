package main

import "fmt"

func main() {
	fmt.Print("Enter #feet to convert to #meters: ")
	var inputFeet float64
	var outputMeters float64

	fmt.Scanf("%f", &inputFeet)

	// convert to meters
	outputMeters = inputFeet * 0.3048

	outputString := fmt.Sprintf("%v feet = %v meters", inputFeet, outputMeters)
	fmt.Println(outputString)
}
