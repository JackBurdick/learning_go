package main

import "fmt"

// Don't fully understand this yet

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	jack := new(int)
	// jack = 4
	fmt.Printf("&jack = %v\n", &jack)
	fmt.Printf("jack = %v\n", jack)
	fmt.Printf("*jack = %v\n", *jack)
	one(jack)
	fmt.Println(*jack)

}
