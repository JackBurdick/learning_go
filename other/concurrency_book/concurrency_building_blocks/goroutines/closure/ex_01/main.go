package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	term := "outide"
	wg.Add(1)
	go func() {
		defer wg.Done()
		term = "inside"
	}()
	wg.Wait()
	fmt.Println(term) // "inside"
}

// this will print "inside"
// goroutines execute within the same address space they were created in.
