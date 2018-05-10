package main

import "fmt"

// functions can return the specified value

func f1() (r int) {
	r = 2
	return
}

func main() {
	jack := f1()
	fmt.Println(jack)
}
