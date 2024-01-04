package main

import (
	"flag"
	"fmt"
	"os"
)

var commands = map[string]string{
	"get":    "get a project",
	"add":    "add a project",
	"delete": "remove a project",
}

type InitFunc func(initType)
type initType int

const (
	A initType = iota
	B
	C
	D
	MaxInitType
)

func Init1(t initType) {
	fmt.Println("Init1 called with type", t)
}

var initFuncs = map[initType]InitFunc{
	A: Init1,
}

func main() {
	// Create map of flags.
	var flagCommands = map[string]map[*flag.FlagSet]string{}

	// Create flags from commands map.
	for k, v := range commands {
		flagCommands[k] = map[*flag.FlagSet]string{
			flag.NewFlagSet(k, flag.ExitOnError): v,
		}
	}

	// Check if no commands were sent.
	if len(os.Args) < 2 {
		rootHelp(commands)
		os.Exit(1)
	}

	// Route to the correct handler based on command.
	for k := range commands {
		if k == os.Args[1] {
			testHandler(os.Args[1])
			os.Exit(0)
		}
	}

	rootHelp(commands)
	os.Exit(1)
}

func testHandler(command string) {

}

func rootHelp(commands map[string]string) {
	fmt.Println(`Projexts is a tool for managing customers.

Usage:
		
	pjx <command> [arguments]
		
	The commands are:
	`)
	// TODO Fix the random order issue:
	// TODO Fix formatting of space between command and description to align:
	// Print all commands and their descriptions.
	for k, v := range commands {
		fmt.Printf("	%v		%v\n", k, v)
	}
	fmt.Printf("\n%v\n", `Use "pjx help <command>" for more information about a command.`)
}
