// package args is to construct and consume environment and command arguments and expose them as
// immutable read-only functions.
package args

import (
	"flag"
	"fmt"
)

var command, managed string

func init() {
	commandHelp := "Specify the command to perform against a managed media folder."
	flag.StringVar(&command, "command", "stats", commandHelp)
	flag.StringVar(&command, "c", "stats", commandHelp+" (shorthand)")

	managedHelp := "Managed folder where media is stored."
	flag.StringVar(&managed, "managed", ".", managedHelp)
	flag.StringVar(&managed, "m", ".", managedHelp+" (shorthand)")
	flag.Parse()
}

type Command int64

const (
	Stats Command = iota
	Add
	Delete
	Tag
)

func (c Command) String() string {
	switch c {
	case Stats:
		return "Stats"
	case Add:
		return "Add"
	case Delete:
		return "Delete"
	case Tag:
		return "Tag"
	}
	panic(fmt.Errorf("Unknown command!"))
}

func GetCommand() Command {
	return Stats // only supported command atm - default for now
}

func GetManagedFolder() string {
	return managed
}
