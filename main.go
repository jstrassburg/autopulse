package main

import (
	"fmt"
	"os"
)

func main() {
	username, password := processArguments()
	fmt.Println("Username:", username, "Password:", password)
}

func processArguments() (string, string) {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Usage:", os.Args[0], "{username} {password}")
		os.Exit(-1)
	}

	return args[1], args[2]
}