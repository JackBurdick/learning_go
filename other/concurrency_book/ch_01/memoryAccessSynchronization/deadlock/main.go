package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()

		// attempt to enter critical section
		v1.mu.Lock()

		// exit critical section before printSum
		defer v1.mu.Unlock()

		// simulate work, and trigger deadlock, this is introduces a race
		// condition and is not a 'perfect' deadlock - which would require
		// synchronization.
		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)

	// locks a then attempts to lock b
	go printSum(&a, &b)

	// locks b, and attempts to lock a
	go printSum(&b, &a)

	wg.Wait()
}
