package main

import "fmt"

func halfEvenBool(val int) (half int, even bool) {
	half = val / 2
	even = (val%2 == 0)
	return
}

func main() {
	for i := 1; i < 10; i++ {
		halfVal, evenBool := halfEvenBool(i)
		fmt.Printf("half(%v) = (%v, %v)\n", i, halfVal, evenBool)
	}

}
