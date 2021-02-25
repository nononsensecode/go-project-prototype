package main

import (
	"fmt"
	"os"

	utils "nononsensecode.com/my_project/internal/utils"
	greetings "nononsensecode.com/my_project/pkg/greetings"
)

func main() {
	args := os.Args[1:]

	var name string
	if utils.ItemExists(args, 0) {
		name = args[0]
	} else {
		name = "Unknown"
	}

	fmt.Printf("%s\n", greetings.Hello(name))
}

