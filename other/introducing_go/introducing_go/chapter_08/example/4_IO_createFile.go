package main

import "os"

func main() {
	file, err := os.Create("test_create.txt")
	if err != nil {
		// handle the error
		return
	}
	defer file.Close()

	// will overwrite current contents
	file.WriteString("WHOOP\thi\nok")
}
