package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
}

func main() {
	x := 1.5
	// changes the value of x 'in place'
	square(&x)
	fmt.Printf("%v\n", x)
}
