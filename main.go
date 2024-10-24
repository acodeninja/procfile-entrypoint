package main

import (
	"log"
	"os"
)

func main() {
	var entrypoint string
	var err error

	args := os.Args[1:]
	if len(args) == 0 {
		entrypoint = "web"
		log.Println("Entrypoint not given, defaulting to 'web'")
	} else {
		entrypoint = args[0]
		log.Printf("Entrypoint given '%s'", entrypoint)
	}

	command, err := getProcfileCommand(entrypoint)

	if err != nil {
		log.Println("Error opening Procfile")
		log.Fatal(err)
	}

	log.Printf("Executing %s: %s", entrypoint, command)

	err = runCommand(command, os.Stdout, os.Stderr)
	if err != nil {
		log.Fatal(err)
	}
}
