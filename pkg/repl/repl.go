package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kyoduke/pokedex/internal/pokeapi"
	"github.com/kyoduke/pokedex/pkg/commands"
)

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	textSlice := strings.Split(trimmedText, " ")
	var cleanedInput []string
	for i, word := range textSlice {
		if word == "" {
			continue
		}
		if i == 0 {
			word = strings.ToLower(word)
		}
		cleanedInput = append(cleanedInput, word)
	}
	return cleanedInput
}

func StartRepl() {
	cfg := commands.Config{
		PokeapiClient:       pokeapi.NewClient(time.Minute, time.Minute*5),
		NextLocationAreaURL: nil,
		PrevLocationAreaURL: nil,
	}

	supportedCommands := commands.GetCommands()

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()
		line := reader.Text()
		cleanedInput := cleanInput(line)
		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]

		if _, ok := supportedCommands[commandName]; !ok {
			fmt.Printf("Command not supported: %s\n", commandName)
			continue
		}

		params := cleanedInput[1:]

		err := supportedCommands[commandName].Callback(&cfg, params)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		if err := reader.Err(); err != nil {
			fmt.Printf("Error during scanning: %v\n", err)
		}
	}

}
