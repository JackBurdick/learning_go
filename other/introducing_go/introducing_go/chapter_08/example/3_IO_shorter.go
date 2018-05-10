package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("test1.txt")
	if err != nil {
		//handle error
		return
	}
	str := string(bs)
	fmt.Println(str)
}
