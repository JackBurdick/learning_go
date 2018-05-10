package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, val := range []string{"a", "b", "c"} {
		wg.Add(1)
		go func(val string) {
			defer wg.Done()
			fmt.Println(val) // c, a, b
		}(val)
	}
	wg.Wait()
}

// now the 'correct'/'expected' value will be printed
