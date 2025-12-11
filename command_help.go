package main

import "fmt"

func commandHelp(cfg *config) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	for k, v := range commands {
		fmt.Printf("%s: %s \n", k, v.description)
	}

	return nil
}
