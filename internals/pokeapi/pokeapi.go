package pokeapi

import "github.com/3h04m1/pokedexcli/internals/pokecache"

type PokeAPI struct {
	LocationArea     LocationArea
	LocationExplorer LocationExplorer
	GetPokemon       GetPokemon
}

func NewPokeAPI(cache *pokecache.Cache) PokeAPI {
	return PokeAPI{
		LocationArea: LocationArea{
			cache: cache,
		},
		LocationExplorer: LocationExplorer{
			cache: cache,
		},
		GetPokemon: GetPokemon{
			cache: cache,
		},
	}
}
