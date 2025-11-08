package commands

import (
	"github.com/kyoduke/pokedex/internal/pokeapi"
)

type Config struct {
	PokeapiClient       pokeapi.PokeapiClient
	NextLocationAreaURL *string
	PrevLocationAreaURL *string
	CatchedPokemons     map[string]pokeapi.RespPokemon
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
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
		"explore": {
			Name:        "explore",
			Description: "Display pokemons in a given location",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Try to catch a pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a pokemon in your pokedex",
			Callback:    commandInspect,
		},
	}
}
