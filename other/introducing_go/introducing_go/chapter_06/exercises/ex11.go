package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	// initialize fake data
	x1, y1 := 3, 7
	// show data
	fmt.Printf("(x1, y1)\t= (%v, %v)\n", x1, y1)
	// swap data
	swap(&x1, &y1)
	fmt.Println("Swap Variables")
	//print swapped data
	fmt.Printf("(x1, y1)\t= (%v, %v)\n", x1, y1)

}
