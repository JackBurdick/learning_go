package main

import "fmt"

func makeOddGenerator() func() uint {
	// start at 1
	i := uint(1)

	// return func; add 2 to i
	//    * i persists between calls
	return func() (returnVal uint) {
		returnVal = i
		i += 2
		return
	}
}

func main() {
	nextOdd := makeOddGenerator()
	// the "i" created persists between function calls
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println("HI")
	fmt.Println(nextOdd())
}
