// io package -> Reader and Writer

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test1.txt")
	if err != nil {
		// handle error
		return
	}
	// make sure file is closed as soon as the function completes
	defer file.Close()

	// get size of file
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
