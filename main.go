package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Initializing REPL as %s\n", user.Username)
	fmt.Printf("***************\n")
	repl.Start(os.Stdin, os.Stdout)
}
