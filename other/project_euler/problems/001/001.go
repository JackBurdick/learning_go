package main

import "fmt"

func main() {
	sum := 0
	RANGEMAX := 1000
	for i := 0; i < RANGEMAX; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("SUM=", sum)
}
