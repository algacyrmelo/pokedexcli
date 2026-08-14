package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locationAreas, err := cfg.pokeapiClient.ListLocationAreas(
		cfg.nextLocationAreasURL,
	)
	if err != nil {
		return err
	}

	cfg.nextLocationAreasURL = locationAreas.Next
	cfg.prevLocationAreasURL = locationAreas.Previous

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationAreasURL == nil {
		return errors.New("You're on the first page")
	}

	locationAreas, err := cfg.pokeapiClient.ListLocationAreas(
		cfg.prevLocationAreasURL,
	)
	if err != nil {
		return err
	}

	cfg.nextLocationAreasURL = locationAreas.Next
	cfg.prevLocationAreasURL = locationAreas.Previous

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
