package main

import (
	"fmt"
	"sync"
)

func sayHi(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hi")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go sayHi(&wg)

	// This is the join point.
	// deterministically block the main goroutine until the goroutine hosting
	// the sayHi() terminates.
	wg.Wait()

	fmt.Println("done")
}

//
