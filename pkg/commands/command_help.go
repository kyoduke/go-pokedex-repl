package commands

import "fmt"

func commandHelp(cfg *Config, args ...string) error {
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
