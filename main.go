package main

import (
	"time"

	"github.com/algacyrmelo/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Minute*5)
	cfg := &config{
		commands: getCommands(),
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
