package handler

import (
	"flag"
	"fmt"
	"os"
)

func deleteRoot() {
	commands := map[string]string{
		"0": "blank",
	}
	fmt.Printf("An option for deleting Projects.\n\nUsage:\n\n	project delete <name> [arguments] \n\nThe arguments are: \n\n")
	for command, description := range commands {
		fmt.Printf("	%v 	%v\n", command, description)
	}
	fmt.Println("")
	os.Exit(1)
}

func deleteHandler(deleteCMD *flag.FlagSet) {
	// Make sure an arg was passed as a name.
	if len(os.Args) < 3 {
		deleteRoot()
		os.Exit(1)
	}

	projectName := os.Args[2]

	fmt.Print("Delete: ")
	fmt.Println(projectName)
}
