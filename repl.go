package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func statRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pockedex > ")

		scanner.Scan()

		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		commands := getCommands()

		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command. Type 'help' for available commands.")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List of next locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List of previouse locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore {area}",
			description: "List of pokemons in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon}",
			description: "Catch pokemon in the area",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect {pokemon}",
			description: "Inspect pokemon from your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Inspect your pokedex",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}
