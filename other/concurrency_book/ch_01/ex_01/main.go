package main

import "fmt"

func main() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("THE VALUE IS %v.\n", data)
	}
	for i := 0; i <= 80; i++ {
		fmt.Printf("%v; the value is %v.\n", i, data)
	}

}
