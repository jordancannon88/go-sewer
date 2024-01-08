package handler

import (
	"fmt"
)

// RootHelp needed to be exported (capitalized)
func RootHelp(commands map[string]string) {
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
