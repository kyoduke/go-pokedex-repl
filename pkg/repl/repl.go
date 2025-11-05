package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

		err := supportedCommands[commandName].Callback()
		if err != nil {
			fmt.Printf("There was an error: %v\n", err)
		}

		if err := reader.Err(); err != nil {
			fmt.Printf("Error during scanning: %v\n", err)
		}
	}

}
