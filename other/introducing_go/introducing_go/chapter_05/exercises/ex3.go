package main

import "fmt"

func main() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	y := x[2:5]
	fmt.Println(y)
}
