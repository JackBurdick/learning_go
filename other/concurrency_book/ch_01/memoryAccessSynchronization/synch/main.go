package main

import (
	"fmt"
	"sync"
)

// The data race is solved, but we still have a race condition.
// we have successfully sychronized access to the memory, but the order of
// operations is still nondeterministic.

func main() {
	var memoryAccess sync.Mutex
	var value int

	go func() {

		// until we declare that until we state otherwise (unlock), our
		// goroutine should have exclusive access to the memory.
		memoryAccess.Lock()
		value++

		// goroutine is done with the memory
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("The value is %v.\n", value)
	}
	memoryAccess.Unlock()
}
