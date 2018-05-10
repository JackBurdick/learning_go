package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")
	if err != nil {
		// handle error
		return
	}
	defer dir.Close()

	// passing in -1 => return all of the entries
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	// will print names of all files in directory
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	// there is also a walk dir

}
