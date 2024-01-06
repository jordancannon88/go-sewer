package handler

import (
	"flag"
	"fmt"
	"os"
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

func getHandler(getCMD *flag.FlagSet, getAll *bool, getID *string) {
	// Make sure an arg was passed as a name.
	if len(os.Args) < 3 {
		rootGet()
		os.Exit(1)
	}

	getCMD.Parse(os.Args[2:])
	projectName := os.Args[2]
	if *getAll == false && *getID == "" {
		fmt.Println("id is required or specify --all for all companies")
		os.Exit(1)
	}
	if *getAll == false && *getID != "" {
		fmt.Print("Get company by ID: ", *getID)
	}
	if *getAll {
		projects := getProject(projectName)
		for _, v := range projects {
			fmt.Printf("%v\n", v.Name)
		}
	}
}

func validateProject(addCMD *flag.FlagSet) {

	// Parse all arguments after the name arg.
	addCMD.Parse(os.Args[3:])

	// if *id == "" || *title == "" || *url == "" || *imageUrl == "" || *description == "" {
	// 	fmt.Print("all fields are required for adding a video")
	// 	addCmd.PrintDefaults()
	// 	os.Exit(1)
	// }

}
