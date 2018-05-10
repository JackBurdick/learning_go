package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// Define flags
	maxp := flag.Int("max", 6, "the max value")

	//Parse
	flag.Parse()

	//Generate a number between 0 and Max
	fmt.Println(rand.Intn(*maxp))
}

/*
	Usage:
		• go run 10_cmdLineArgs.go
		• go run 10_cmdLineArgs.go -max=100

*/
