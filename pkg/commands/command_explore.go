package commands

import "fmt"

func commandExplore(c *Config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("cannot call explore without an area name\nExample: explore canalave-city-area\n")
	}
	fmt.Println("called explore for location", params[0])
	return nil
}
