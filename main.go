package main

import (
	"fmt"
	"os"
)

func main() {
	name := "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	
	fmt.Printf("Hello, %s! Welcome to the Secure Supply Chain Demo.\n", name)
}
