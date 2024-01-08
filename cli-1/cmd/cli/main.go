package main

import (
	"flag"
	"fmt"
	"internal/handler"
	"os"
)

var commands = map[string]string{
	"get":    "get a project",
	"add":    "add a project",
	"delete": "remove a project",
}

// var commands1 = map[string]func(){
// 	"get": func() {},
// }

// rootHelp(commands)

func main() {
	// Create map of flags.
	var flagCommands = map[string]map[*flag.FlagSet]string{}

	// Create flags from commands map.
	for k, v := range commands {
		flagCommands[k] = map[*flag.FlagSet]string{
			flag.NewFlagSet(k, flag.ExitOnError): v,
		}
	}

	// Create map of functions.
	// var functionCommands = map[string]func(){}

	// Add functions to functionCommands map.
	// for k := range commands {
	// 	commands1[k] = test
	// }

	// Check if no commands were sent.
	if len(os.Args) < 2 {
		// rootHelp(commands)
		os.Exit(1)
	}

	// Route to the correct handler based on command.
	// for k := range commands {
	// 	if k == os.Args[1] {
	// 		f := functionCommands[k]
	// 		f()
	// 		os.Exit(0)
	// 	}
	// }

	switch os.Args[1] {
	case "help":
		handler.RootHelp(commands)
		// rootHelp(commands)
	case "get":
		// getHandler(getCMD, getAll, getID)
	case "add":
		// addHandler(addCMD, addDesc)
	case "delete":
		// deleteHandler(deleteCMD)
	default:
		subcommand := os.Args[1]
		fmt.Printf("Unknown subcommand: %s\nRun 'project help' for usage.\n", subcommand)
		os.Exit(1)
	}

	// rootHelp(commands)
	os.Exit(1)
}
