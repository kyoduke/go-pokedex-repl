package commands

import "fmt"

func commandMap(cfg *Config, args ...string) error {
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

func commandMapBack(cfg *Config, args ...string) error {
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
