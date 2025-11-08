package commands

import (
	"fmt"
	"maps"
)

func commandPokedex(c *Config, args ...string) error {
	pokemons := maps.Keys(c.CatchedPokemons)
	fmt.Println("Your pokedex:")
	for pokemon := range pokemons {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}
