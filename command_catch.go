package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, arg ...string) error {

	if len(arg) == 0 {
		fmt.Println("Please provide the name of Pokemon.")
	}

	fmt.Println("Throwing a Pokeball at " + arg[0] + "...")

	pokemon, err := cfg.pokeapiClient.PokemonDetails(arg[0])
	if err != nil {
		return err
	}

	maxBaseExp := 300
	randomNumber := rand.Intn(100)
	chance := 100 - (pokemon.BaseExperience * 100 / maxBaseExp)

	if randomNumber < chance {
		fmt.Println("🎉 Pokemon caught!")
		cfg.caughtPokemon[arg[0]] = pokemon
	} else {
		fmt.Println("💨 Pokemon escaped!")
		return nil
	}
	return nil
}
