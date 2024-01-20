package storage

type Storage struct {
	Pokemons PokemonStorage
}

func NewStorage() Storage {
	return Storage{
		Pokemons: NewPokemonStorage(),
	}
}
