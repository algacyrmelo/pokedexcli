package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const captureThreshold int = 50
const minEscapeMargin int = 10

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	// Prevent 100% chance of capture. The roll range is never smaller than
	// the threshold plus a small margin.
	rollRange := max(pokemon.BaseExperience, captureThreshold+minEscapeMargin)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	// Randomly choose a number from 0 to the rollRange.
	roll := rand.Intn(rollRange)

	// The higher the base experience, the less likely the capture.
	if roll < captureThreshold {
		cfg.caughtPokemon[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
