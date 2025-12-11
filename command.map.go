package main

import (
	"fmt"

	"github.com/mati99789/pokedexcli/internal/pokeapi"
)

var defaultUrl = "https://pokeapi.co/api/v2/location-area?limit=20"

func commandMap(cfg *config, args ...string) error {

	if cfg.next == nil {
		cfg.next = &defaultUrl

	}
	result, err := cfg.pokeapiClient.FetchLocationAreas(*cfg.next)

	if err != nil {
		return err
	}

	cfg.next = result.Next
	cfg.previous = result.Previous

	printResultPokeApi(result)
	return nil
}

func printResultPokeApi(result pokeapi.PokeAPI) {
	for _, v := range result.Results {
		fmt.Println(v.Name)
	}
}

func commandMapB(cfg *config, args ...string) error {

	if cfg.previous == nil {
		fmt.Println("You're on the first page")

		return nil
	}

	result, err := cfg.pokeapiClient.FetchLocationAreas(*cfg.previous)

	if err != nil {
		return err
	}

	cfg.next = result.Next
	cfg.previous = result.Previous

	printResultPokeApi(result)
	return nil
}
