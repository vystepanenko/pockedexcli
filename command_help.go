package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pockedex CLI!")
	fmt.Println("Available commands:")

	for _, command := range getCommands() {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}

	fmt.Println("")

	return nil
}
