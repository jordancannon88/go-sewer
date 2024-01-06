package project

import (
	"encoding/json"
	"os"
)

type project struct {
	Name        string
	Description string
}

func getProject(projectName string) (project []project) {

	projectByte, err := os.ReadFile("./project.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(projectByte, &project)

	if err != nil {
		panic(err)
	}

	return project
}

func saveProject(project []project) {

	projectByte, err := json.Marshal(project)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("./project.json", projectByte, 0644); err != nil {
		// log.Fatal(err)
		panic(err)
	}
}
