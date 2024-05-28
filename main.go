package main

import (
	"time"

	"github.com/adspampin/pokedexcli/internal/pokeapi"
)

func main() {
    pokeClient := pokeapi.NewClient(5 * time.Second, 10 * time.Second)
    cfg := &config{
        pokeapiClient: pokeClient,
    }
    startRepl(cfg)
}
