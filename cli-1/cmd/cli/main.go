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

func main() {
	// Check if no subcommmands were sent.
	if len(os.Args) < 2 {
		handler.RootHelp(commands)
		os.Exit(1)
	}

	// Create subcommand map.
	var subCommand = map[string]*flag.FlagSet{}

	createSubCommands(subCommand, commands)

	// Create arguments.
	getCMD := subCommand["get"]
	getAll := getCMD.Bool("all", false, "Test")
	getID := getCMD.String("id", "", "Test")

	addCMD := subCommand["add"]
	addDesc := addCMD.String("desc", "", "Test")

	// TODO Convert this to a dynamic function.
	switch os.Args[1] {
	case "help":
		handler.RootHelp(commands)
	case "get":
		handler.GetHandler(getCMD, getAll, getID)
	case "add":
		handler.AddHandler(addCMD, addDesc)
	case "delete":
		// deleteHandler(deleteCMD)
	default:
		subcommand := os.Args[1]
		fmt.Printf("Unknown subcommand: %s\nRun 'project help' for usage.\n", subcommand)
		os.Exit(1)
	}
}

func createSubCommands(flagCommands map[string]*flag.FlagSet, customCommands map[string]string) map[string]*flag.FlagSet {
	for k := range customCommands {
		flagCommands[k] = flag.NewFlagSet(k, flag.ExitOnError)
	}

	return flagCommands
}
