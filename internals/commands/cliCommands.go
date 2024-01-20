package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/3h04m1/pokedexcli/internals/pokeapi"
	"github.com/3h04m1/pokedexcli/internals/storage"
)

// CliCommand is a struct that represents a command that can be run in the CLI.
// It has a name, a description, and a callback function.
// The callback function is executed when the command is run.

type CommandCallback func(args ...string) error

type CliCommand struct {
	Name        string
	Description string
	Callback    CommandCallback
}

// CliCommands is a struct that represents a collection of commands.
// It has a prefix, an error prefix, and a list of commands.
// The prefix is displayed before the user input.
// The error prefix is displayed before an error message.
// The list of commands is a list of CliCommand structs.
type CliCommands struct {
	api         pokeapi.PokeAPI
	prefix      string
	errorPrefix string
	storage     storage.Storage
	commands    []CliCommand
}

// respond is a helper function that prints a response to the CLI.
// It prints the response in green text.
func (c *CliCommands) respond(response any) {
	fmt.Printf("\033[32m%+v\033[0m\n", response)
}

func (c *CliCommands) RegisterCommand(commandFunc func(*CliCommands) CliCommand) {
	c.commands = append(c.commands, commandFunc(c))
}

// Run runs the Pokedex.
func (c *CliCommands) Run() {
	reader := bufio.NewReader(os.Stdin)
	c.respond("Welcome to the Pokedex")
	c.respond("Type `help` for a list of commands")

	for {
		fmt.Print(c.prefix)
		var input string
		input, _ = reader.ReadString('\n')
		wordList := strings.Split(strings.Trim(input, "\n"), " ")

		commandName := wordList[0]
		args := wordList[1:]
		command, found := c.getCommand(commandName)
		if !found {
			message := fmt.Errorf("%scommand `%s` not found", c.errorPrefix, input)
			fmt.Println(message)
			continue
		}
		err := command.Callback(args...)
		if err != nil {
			message := fmt.Errorf("%serror executing command `%s`: \n\t%s", c.errorPrefix, command.Name, err.Error())
			fmt.Println(message)
		}
	}
}

// It gets a command by name.
func (c *CliCommands) getCommand(input string) (CliCommand, bool) {
	for _, command := range c.commands {
		if command.Name == input {
			return command, true
		}
	}
	return CliCommand{}, false
}
