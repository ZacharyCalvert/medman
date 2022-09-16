package main

import (
	"fmt"

	"github.com/ZacharyCalvert/medman/args"
)

func main() {
	fmt.Printf("Hello world will run %s command\n", args.GetCommand())
}
