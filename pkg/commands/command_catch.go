package commands

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(c *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("catch command should not be nil")
	}
	pokemon := args[0]

	if _, ok := c.CatchedPokemons[pokemon]; ok {
		fmt.Printf("%s is already in your pokedex\n", pokemon)
		return nil
	}

	respPokemon, err := c.PokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	catched := catchAttempt(respPokemon.BaseExperience)
	if !catched {
		fmt.Printf("Failed to catch %s\n", pokemon)
		return nil
	}

	c.CatchedPokemons[args[0]] = respPokemon
	fmt.Printf("Successfully captured %s\n", pokemon)

	return nil
}

func catchAttempt(baseExperience int) bool {
	const MIDPOINT_EXP = 150
	const DIFFICULTY_FACTOR = 0.05
	exponent := DIFFICULTY_FACTOR * (float64(baseExperience) - float64(MIDPOINT_EXP))
	probability := 1 / (1 + math.Exp(exponent))
	randomRoll := rand.Float32()
	return randomRoll <= float32(probability)
}
