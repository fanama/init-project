package main

import (
	"fanama/init-project/infra/gitInfra"
	"fmt"
	"log"
)

func main() {
	fmt.Println("=== project builder ===")

	var project string

	fmt.Print("Name of the project : ")
	fmt.Scanln(&project)

	fmt.Println("Creating the project")
	err := gitInfra.Clone(project)

	if err != nil {
		log.Fatal(err)
	}

	err = gitInfra.RemoveOrigin(project)

	if err != nil {
		log.Fatal(err)
	}

}
