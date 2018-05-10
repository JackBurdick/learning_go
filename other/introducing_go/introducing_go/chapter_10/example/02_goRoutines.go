package main

import (
	"fmt"
	"math/rand"
	"time"
)

func func_ex(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		sleepDuration := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * sleepDuration)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go func_ex(i)
	}
	var input string
	fmt.Scanln(&input)
}

// output shows that goroutines are run simultaneously, not sequentially
