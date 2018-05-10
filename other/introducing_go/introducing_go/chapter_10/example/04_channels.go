package main

import (
	"fmt"
	"time"
)

// send "ping" to channel c
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// send "ping" to channel c
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// print message received from channel c
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// make channel
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

// NOTE:
// will take turns printing ping and pong
