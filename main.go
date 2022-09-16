package main

import (
	"github.com/ZacharyCalvert/medman/args"
	"github.com/ZacharyCalvert/medman/stat"
)

func main() {
	command := args.GetCommand()
	switch command {
	case args.Stats:
		stat.RunCommand()
	default:
		panic("Command not implemented")
	}
}
