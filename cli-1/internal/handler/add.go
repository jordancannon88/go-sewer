package handler

import (
	"flag"
	"fmt"
	"os"

	"internal/project"
)

func rootAdd() {
	commands := map[string]string{
		"-desc": "a description of the project",
	}
	fmt.Printf("An option for adding Projects.\n\nUsage:\n\n	project add <name> [arguments] \n\nThe arguments are: \n\n")
	for command, description := range commands {
		fmt.Printf("	%v 	%v\n", command, description)
	}
	fmt.Println("")
	os.Exit(1)
}

func AddHandler(addCMD *flag.FlagSet, addDesc *string) {
	// Make sure an arg was passed as a name.
	if len(os.Args) < 3 {
		rootAdd()
		os.Exit(1)
	}

	addCMD.Parse(os.Args[3:])

	projectName := os.Args[2]

	proj := project.Project{
		Name:        projectName,
		Description: *addDesc,
	}

	fmt.Println(proj)

	// Get all projects.
	projects := project.Get()

	// Check if project name already exists.
	for _, v := range projects {
		if v.Name == projectName {
			fmt.Printf("Project name already exists: %v\n", projectName)
			os.Exit(1)
		}
	}

	// Add new project to all projects.
	projects = append(projects, proj)

	// Save all projects.
	project.Save(projects)
}
