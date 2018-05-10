package main

import "fmt"

// defer is usful for when resources needed to be freed in some way
// When we open a file --> close it later
// f, _ := os.Open(filename)
// defer f.Close()
//
// ----- advantage -----
// • readability/"understandablity": keep close call near open
// • Deferred functions are run even if a runtime panic occurs

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer second()
	first()
}
