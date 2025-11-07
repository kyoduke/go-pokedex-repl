package commands

import (
	"fmt"
	"os"

	"github.com/kyoduke/pokedex/internal/pokeapi"
)

type Config struct {
	PokeapiClient       pokeapi.PokeapiClient
	NextLocationAreaURL *string
	PrevLocationAreaURL *string
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the program",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Display how to use the program",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Display a list of maps",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display previous maps",
			Callback:    commandMapBack,
		},
	}
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *Config) error {
	locationResp, err := cfg.PokeapiClient.ListLocationAreas(cfg.NextLocationAreaURL)
	if err != nil {
		return fmt.Errorf("[ERROR] Error when using map command %w", err)
	}
	cfg.NextLocationAreaURL = locationResp.Next
	cfg.PrevLocationAreaURL = locationResp.Previous
	areas := locationResp.Results
	fmt.Println()
	for _, area := range areas {
		fmt.Printf("- %s\n", area.Name)
	}
	return nil
}

func commandMapBack(cfg *Config) error {
	if cfg.PrevLocationAreaURL == nil {
		fmt.Println("You are on the first page")
		return nil
	}

	locationResp, err := cfg.PokeapiClient.ListLocationAreas(cfg.PrevLocationAreaURL)
	if err != nil {
		return fmt.Errorf("[ERROR] Error when using mapb command %w", err)
	}

	cfg.NextLocationAreaURL = locationResp.Next
	cfg.PrevLocationAreaURL = locationResp.Previous
	for _, area := range locationResp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	return nil
}
