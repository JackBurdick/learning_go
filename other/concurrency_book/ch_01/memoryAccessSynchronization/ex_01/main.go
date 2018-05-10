package main

import (
	"fmt"
)

func main() {

	/*
	  critical section: section of the program that needs exclusive access to
	  a shared resource. There are three in this example.
	*/

	var data int

	// critical section 1
	go func() {
		data++
	}()

	// critical section 2
	if data == 0 {
		fmt.Println("the value is 0.")
	} else {

		// critical section 3
		fmt.Printf("The value is %v.\n", data)
	}

}
