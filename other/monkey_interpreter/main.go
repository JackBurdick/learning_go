package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/JackBurdick/monkey_interpreter/repl"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Let's write some monkey\n", usr.Username)
	fmt.Printf("Enter commands below\n")
	repl.Start(os.Stdin, os.Stdout)
}
