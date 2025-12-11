package main

import (
	"time"

	"github.com/mati99789/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
}

func main() {
	pokeapi := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := config{
		pokeapiClient: pokeapi,
	}
	startRepl(&cfg)
}
