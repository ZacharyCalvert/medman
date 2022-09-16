// package args is to construct and consume environment and command arguments and expose them as
// immutable read-only functions.
package args

import (
	"flag"
	"fmt"
)

var command, managed string

/*
Simply load up all of the flags on the init of the package.  One time, only time.
Safe to assume this pacakge will be loaded early as it dictates the entire behavior
of the application.
*/
func init() {
	commandHelp := "Specify the command to perform against a managed media folder."
	flag.StringVar(&command, "command", "stats", commandHelp)
	flag.StringVar(&command, "c", "stats", commandHelp+" (shorthand)")

	managedHelp := "Managed folder where media is stored."
	flag.StringVar(&managed, "managed", ".", managedHelp)
	flag.StringVar(&managed, "m", ".", managedHelp+" (shorthand)")
	flag.Parse()
}

// Command represents the command the application is being instructed to perform
type Command int64

const (
	Stats Command = iota
	Add
	Delete
	Tag
)

// Stringer adherence
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

// GetCommand returns the enumerated (iota) representation of the requested command
func GetCommand() Command {
	return Stats // only supported command atm - default for now
}

// GetManagedFolder returns the defined managed folder which has NOT been validated for correctness
func GetManagedFolder() string {
	return managed
}
