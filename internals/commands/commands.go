package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/3h04m1/pokedexcli/internals/pokeapi"
	"github.com/3h04m1/pokedexcli/internals/storage"
	"github.com/3h04m1/pokedexcli/internals/utils"
)

func getHelpCommand(commands *CliCommands) CliCommand {
	return CliCommand{
		Name:        "help",
		Description: "Displays help information",
		Callback: func(args ...string) error {
			for _, command := range commands.commands {
				commands.respond(fmt.Sprintf("%s - %s", command.Name, command.Description))
			}
			return nil
		},
	}
}

func getExitCommand(commands *CliCommands) CliCommand {
	return CliCommand{
		Name:        "exit",
		Description: "Exits the Pokedex",
		Callback: func(args ...string) error {
			commands.respond("Goodbye")
			os.Exit(0)
			return nil
		},
	}
}

func getMapCommand(commands *CliCommands) CliCommand {
	return CliCommand{
		Name:        "map",
		Description: "Displays the next names of 20 location areas",
		Callback: func(args ...string) error {
			requestResponse, err := commands.api.LocationArea.Next()
			if err != nil {
				return errors.New("error fetching location areas")
			}
			for _, result := range requestResponse {
				commands.respond(result.Name)
			}
			return nil
		},
	}
}

func getMapBCommand(commands *CliCommands) CliCommand {
	return CliCommand{
		Name:        "mapb",
		Description: "Displays the previous names of 20 location areas",
		Callback: func(args ...string) error {
			requestResponse, err := commands.api.LocationArea.Previous()
			if err != nil {
				return err
			}
			for _, result := range requestResponse.Results {
				commands.respond(result.Name)
			}
			return nil
		},
	}
}

func getExploreCommand(commands *CliCommands) CliCommand {
	return CliCommand{
		Name:        "explore",
		Description: "Explores a location area",
		Callback: func(args ...string) error {
			if len(args) < 1 {
				return errors.New("please provide a location area name")
			}
			location := args[0]
			fmt.Println("Exploring location area ...", location)
			requestResponse, err := commands.api.LocationExplorer.GetLocationAreaEncounters(location)
			if err != nil {
				return err
			}
			fmt.Println("Found Pokemon:")
			for _, result := range requestResponse {
				msg := fmt.Sprintf("  - %s", result)
				commands.respond(msg)
			}
			return nil
		},
	}
}

func getCatchCommand(commands *CliCommands) CliCommand {
	return CliCommand{
		Name:        "catch",
		Description: "Catches a Pokemon",
		Callback: func(args ...string) error {
			if len(args) < 1 {
				return errors.New("please provide a Pokemon name")
			}
			pokemonName := args[0]
			if commands.storage.Pokemons.HasPokemon(pokemonName) {
				commands.respond("You already have this Pokemon!")
				return nil
			}
			fmt.Println("Catching Pokemon ", pokemonName)
			pokemon, err := commands.api.GetPokemon.Get(pokemonName)
			if err != nil {
				return err
			}
			isCatched := utils.RandomBool(pokemon.BaseExperience)
			if isCatched {
				commands.storage.Pokemons.AddPokemon(pokemon)
				commands.respond("Pokemon catched!")
			} else {
				commands.respond("Pokemon escaped!")
			}
			return nil
		},
	}
}

func getInspectCommand(commands *CliCommands) CliCommand {

	return CliCommand{
		Name:        "inspect",
		Description: "Inspects a Pokemon",
		Callback: func(args ...string) error {
			if len(args) < 1 {
				return errors.New("please provide a Pokemon name")
			}
			pokemonName := args[0]
			if !commands.storage.Pokemons.HasPokemon(pokemonName) {
				commands.respond("You don't have this Pokemon!")
				return nil
			}
			fmt.Println("Inspecting Pokemon ", pokemonName)
			pokemon, err := commands.storage.Pokemons.GetPokemon(pokemonName)
			if err != nil {
				return errors.New("you have not caught that pokemon")
			}
			var statsRes string
			for _, stat := range pokemon.Stats {
				statsRes += fmt.Sprintf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
			}
			var typesRes string
			for _, pokemonType := range pokemon.Types {
				typesRes += fmt.Sprintf("  - %s\n", pokemonType.Type.Name)
			}
			response := fmt.Sprintf(`
Name: %s
Base Experience: %d
Height: %d
Weight: %d
Stats:
%s
Types:
%s
			`, pokemon.Name, pokemon.BaseExperience, pokemon.Height, pokemon.Weight, statsRes, typesRes)
			commands.respond(response)
			return nil
		},
	}
}

func getPokedexCommand(commands *CliCommands) CliCommand {
	return CliCommand{
		Name:        "pokedex",
		Description: "Displays your Pokedex",
		Callback: func(args ...string) error {
			pokemons := commands.storage.Pokemons.GetPokemons()
			if len(pokemons) == 0 {
				commands.respond("You have not caught any Pokemon!")
				return nil
			}
			commands.respond("Your Pokedex:")
			for _, pokemon := range pokemons {
				commands.respond("  - " + pokemon)
			}
			return nil
		},
	}
}

func NewCommands(api pokeapi.PokeAPI, storage storage.Storage) CliCommands {
	commands := CliCommands{
		api:         api,
		prefix:      "\033[34mPokedex > \033[0m",
		errorPrefix: "\033[31mPokedex > \033[0m",
		commands:    []CliCommand{},
		storage:     storage,
	}

	commands.RegisterCommand(getHelpCommand)
	commands.RegisterCommand(getExitCommand)
	commands.RegisterCommand(getMapCommand)
	commands.RegisterCommand(getMapBCommand)
	commands.RegisterCommand(getExploreCommand)
	commands.RegisterCommand(getCatchCommand)
	commands.RegisterCommand(getInspectCommand)
	commands.RegisterCommand(getPokedexCommand)
	return commands
}
