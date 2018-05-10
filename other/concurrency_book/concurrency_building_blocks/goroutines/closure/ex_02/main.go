package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, val := range []string{"a", "b", "c"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(val) // c, c, c
		}()
	}
	wg.Wait()
}

// undetermined output.
// `val` memory will be transferred to the heap since the go runtime recognizes
// that it is being referenced.
