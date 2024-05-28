package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("no location area provided")
    }

    locationAreaName := args[0]

    location, err := cfg.pokeapiClient.ListLocation(locationAreaName)
    if err != nil{
        return err
    }

    fmt.Println()
    fmt.Printf("All the pokemon in %v:\n", locationAreaName)
    for _, pokemon := range location.PokemonEncounters {
        fmt.Printf("  - %v\n", pokemon.Pokemon.Name)
    }
    fmt.Println()


    return nil
}
