package app

import (
	"time"

	"github.com/3h04m1/pokedexcli/internals/commands"
	"github.com/3h04m1/pokedexcli/internals/pokeapi"
	"github.com/3h04m1/pokedexcli/internals/pokecache"
	"github.com/3h04m1/pokedexcli/internals/storage"
)

type App struct {
	api            pokeapi.PokeAPI
	commandHandler commands.CliCommands
	storage        storage.Storage
}

func NewApp() App {
	cache := pokecache.NewCache(time.Duration(5) * time.Second)
	api := pokeapi.NewPokeAPI(cache)
	storage := storage.NewStorage()
	commandHandler := commands.NewCommands(api, storage)
	return App{
		api:            api,
		commandHandler: commandHandler,
		storage:        storage,
	}
}

func (a *App) Run() {
	a.commandHandler.Run()
}
