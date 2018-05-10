package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}
func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids := []Person{
		{"Abby", 25},
		{"Jack", 21},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)

	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}

// the sort interface requires three methods;
// • Len
// 		• Return the length of the thing we are sorting
// • Less
// 		• Determine whether the item at position i is strictly less than the item at pos j
// • Swap
// 		• Swap the items
