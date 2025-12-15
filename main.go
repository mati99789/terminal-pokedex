package main

import (
	"time"

	"github.com/mati99789/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func main() {
	pokeAPI := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := config{
		pokeapiClient: pokeAPI,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	startRepl(&cfg)
}
