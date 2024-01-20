package pokeapi

import (
	"github.com/3h04m1/pokedexcli/internals/pokecache"
	"github.com/3h04m1/pokedexcli/internals/utils"
)

type GetLocationAreaRes struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type LocationExplorer struct {
	cache *pokecache.Cache
}

func (l *LocationExplorer) GetLocationArea(name string) (GetLocationAreaRes, error) {
	res, err := utils.Request[GetLocationAreaRes]("https://pokeapi.co/api/v2/location-area/"+name, l.cache)
	if err != nil {
		return GetLocationAreaRes{}, err
	}
	return res, nil

}

func (l *LocationExplorer) GetLocationAreaEncounters(name string) ([]string, error) {
	var encounters []string
	res, err := l.GetLocationArea(name)
	if err != nil {
		return encounters, nil
	}
	for _, encounter := range res.PokemonEncounters {
		pokemonName := encounter.Pokemon.Name
		encounters = append(encounters, pokemonName)
	}

	return encounters, nil
}
