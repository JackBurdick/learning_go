package main

import "fmt"

func makeOddGenerator_trail() func() uint {
	i := uint(1)
	return func() (returnVal uint) {
		returnVal = i
		i += 2
		return
	}
}

func main() {
	next_odd_one := makeOddGenerator_trail()
	next_odd_two := makeOddGenerator_trail()
	// starts at 1
	fmt.Println(next_odd_one())
	fmt.Println(next_odd_one())
	// starts at 1 again
	fmt.Println(next_odd_two())
	fmt.Println(next_odd_one())
}
