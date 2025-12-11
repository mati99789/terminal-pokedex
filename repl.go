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
	callback    func(cfg *config) error
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

		cmd, ok := commands[firstWord]

		if ok {
			cmd.callback(cfg)
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
	}
}
