package handler

import (
	"flag"
	"fmt"
	"os"

	"internal/project"
)

func rootGet() {
	commands := map[string]string{
		"-id":  "get a project by its name",
		"-all": "get all projects",
	}
	fmt.Printf("An option for getting Projects.\n\nUsage:\n\n	project get <name> [arguments] \n\nThe arguments are: \n\n")
	for command, description := range commands {
		fmt.Printf("	%v 	%v\n", command, description)
	}
	fmt.Println("")
	os.Exit(1)
}

func GetHandler(getCMD *flag.FlagSet, getAll *bool, getID *string) {
	// Make sure an arg was passed as a name.
	if len(os.Args) < 3 {
		rootGet()
		os.Exit(1)
	}

	getCMD.Parse(os.Args[3:])

	fmt.Println(*getAll)
	fmt.Println(*getID)

	if *getAll == false && *getID == "" {
		fmt.Println("id is required or specify --all for all companies")
		os.Exit(1)
	}
	if *getAll == false && *getID != "" {
		fmt.Print("Get company by ID: ", *getID)
	}
	if *getAll {
		projects := project.Get()
		for _, v := range projects {
			fmt.Printf("%v\n", v.Name)
		}
	}
}
