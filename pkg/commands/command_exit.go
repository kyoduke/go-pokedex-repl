package commands

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
