package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide the name or ID, eg: explore pastoria-city-area")
		return nil
	}

	locationName := args[0]

	fmt.Println("Exploring pastoria-city-area...")
	result, err := cfg.pokeapiClient.FetchExplore(locationName)

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, v := range result.PokemonEncounters {
		fmt.Printf("%v\n", v.Pokemon.Name)
	}
	return nil
}
