package main

import "fmt"

func zero_val(x int) {
	x = 0
}

func zero_ref(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 53243
	fmt.Printf("x = %v\n", x)
	fmt.Println("Pass by VALUE --> remain unchanged")
	zero_val(x)
	fmt.Printf("x = %v\n", x)
	fmt.Println("Pass by REFERENCE --> modifies variable")
	zero_ref(&x)
	fmt.Printf("x = %v\n", x)
}
