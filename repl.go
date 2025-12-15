package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, arg ...string) error
}

var commands map[string]cliCommand

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		userInputs := cleanInput(scanner.Text())
		firstWord := ""

		if len(userInputs) > 0 {
			firstWord = userInputs[0]
		}

		args := userInputs[1:]
		cmd, ok := commands[firstWord]

		if ok {
			error := cmd.callback(cfg, args...)

			if error != nil {
				fmt.Println(error)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func init() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "go to next list of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "go back to previous list of locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Locations that can be visited within the games. Locations make up sizable portions of regions, like cities or routes.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			callback:    commandInspect,
		},
	}
}
