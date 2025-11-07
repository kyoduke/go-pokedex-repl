package commands

import "fmt"

func commandExplore(c *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("cannot call explore without an area name\nExample: explore canalave-city-area\n")
	}

	res, err := c.PokeapiClient.ListAreaEncounters(&args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", args[0])
	for _, encounter := range res.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
