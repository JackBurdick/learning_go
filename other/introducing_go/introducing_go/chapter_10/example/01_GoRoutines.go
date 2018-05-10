package main

import "fmt"

func func_example(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go func_example(3)
	var input string
	fmt.Scanln(&input)
}
