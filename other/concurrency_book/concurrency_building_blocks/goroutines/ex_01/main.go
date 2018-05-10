package main

import "fmt"

func sayHi() {
	fmt.Println("hi")
}

func main() {
	go sayHi()
	fmt.Println("done")
}

// we don't know whether sayHi() will run before the program exits or not
// could add a sleep before exiting, but this still doesn't guarantee it will run
