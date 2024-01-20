package pokeapi

import (
	"errors"

	"github.com/3h04m1/pokedexcli/internals/pokecache"
	"github.com/3h04m1/pokedexcli/internals/utils"
)

type LocationArea struct {
	cache    *pokecache.Cache
	prevPage string
	nextPage string
}

type LocationAreaReturn struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationAreaRes utils.PaginatedResponse[LocationAreaReturn]

const baseURL = "https://pokeapi.co/api/v2/location-area"

func (l *LocationArea) updatePages(r locationAreaRes) {
	l.prevPage = r.Previous
	l.nextPage = r.Next
}

func (l *LocationArea) get(url string) (locationAreaRes, error) {
	return utils.Request[locationAreaRes](url, l.cache)
}

func (l *LocationArea) Next() ([]LocationAreaReturn, error) {
	if l.nextPage == "" && l.prevPage == "" {
		res, err := l.get(baseURL)
		if err != nil {
			return []LocationAreaReturn{}, err
		}
		l.updatePages(res)
		return res.Results, nil
	}
	res, err := utils.Request[locationAreaRes](l.nextPage, l.cache)
	if err != nil {
		return []LocationAreaReturn{}, err
	}
	l.updatePages(res)
	return res.Results, err
}

func (l *LocationArea) Previous() (locationAreaRes, error) {
	if l.prevPage == "" {
		return locationAreaRes{}, errors.New("You are on the first page")
	}
	res, err := utils.Request[locationAreaRes](l.prevPage, l.cache)
	if err != nil {
		return locationAreaRes{}, err
	}
	l.updatePages(res)
	return res, nil
}
