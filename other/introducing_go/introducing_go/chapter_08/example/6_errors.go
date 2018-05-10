package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Error Message")
	fmt.Println(err)
}
