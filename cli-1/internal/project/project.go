package project

import (
	"encoding/json"
	"os"
)

type Project struct {
	Name        string
	Description string
}

func Get() (project []Project) {

	projectByte, err := os.ReadFile("../../project.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(projectByte, &project)

	if err != nil {
		panic(err)
	}

	return project
}

// takes in any number of projects (could take in none, too, with the spread operator or whatever the hell it's called)
// func Save(project ...Project) {
func Save(project []Project) {

	projectByte, err := json.Marshal(project)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("../../project.json", projectByte, 0644); err != nil {
		// log.Fatal(err)
		panic(err)
	}
}
