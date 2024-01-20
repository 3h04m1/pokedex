package storage

import (
	"errors"

	"github.com/3h04m1/pokedexcli/internals/pokeapi"
)

type PokemonStorage struct {
	Pokedex map[string]pokeapi.GetPokemonRes
}

func NewPokemonStorage() PokemonStorage {
	return PokemonStorage{
		Pokedex: make(map[string]pokeapi.GetPokemonRes),
	}
}

func (s *PokemonStorage) AddPokemon(pokemon pokeapi.GetPokemonRes) {
	s.Pokedex[pokemon.Name] = pokemon
}

func (s *PokemonStorage) GetPokemons() []string {
	res := make([]string, 0)
	for _, pokemon := range s.Pokedex {
		res = append(res, pokemon.Name)
	}
	return res
}

func (s *PokemonStorage) GetPokemon(name string) (pokeapi.GetPokemonRes, error) {
	if s.HasPokemon(name) {
		return s.Pokedex[name], nil
	}
	return pokeapi.GetPokemonRes{}, errors.New("pokemon not found")
}

func (s *PokemonStorage) HasPokemon(name string) bool {
	_, ok := s.Pokedex[name]
	return ok
}

func (s *PokemonStorage) RemovePokemon(name string) error {
	if s.HasPokemon(name) {
		delete(s.Pokedex, name)
		return nil
	}
	return errors.New("pokemon not found")
}
