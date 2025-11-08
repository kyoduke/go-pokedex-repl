package commands

import "fmt"

func commandInspect(c *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Not enough arguments\n")
	}
	pokemon := args[0]

	if val, ok := c.CatchedPokemons[pokemon]; ok {
		var hp, attack, defense, special_attack, special_defense, speed int
		var types string
		for _, val := range val.Stats {
			switch val.Stat.Name {
			case "hp":
				hp = val.BaseStat
			case "attack":
				attack = val.BaseStat
			case "defense":
				defense = val.BaseStat
			case "special-attack":
				special_attack = val.BaseStat
			case "special-defense":
				special_defense = val.BaseStat
			case "speed":
				speed = val.BaseStat
			}
		}

		for _, val := range val.Types {
			types = types + fmt.Sprintf("	- %s\n", val.Type.Name)
		}
		fmt.Printf(`
Name: %v
Height: %v
Weight: %v
Stats:
	-hp: %v
	-attack: %v
	-defense: %v
	-special-attack: %v
	-special-defense: %v
	-speed: %v
Types:
%v
`,
			val.Name,
			val.Height,
			val.Weight,
			hp,
			attack,
			defense,
			special_attack,
			special_defense,
			speed,
			types,
		)
	} else {
		fmt.Println("catch the pokemon first!")
	}
	return nil
}
