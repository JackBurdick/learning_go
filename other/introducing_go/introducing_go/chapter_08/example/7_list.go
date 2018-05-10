package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	// follow list and print each element
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
