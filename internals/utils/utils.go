package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/3h04m1/pokedexcli/internals/pokecache"
)

type AnyStruct[T any] struct{}

type Response[T any] struct {
	Results []T `json:"results"`
}

type JsonRes[T any] struct {
	json T
	raw  []byte
}

type PaginatedResponse[T any] struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

func Request[T any](url string, cache *pokecache.Cache) (T, error) {
	cached, ok := cache.Get(url)
	if !ok {
		res, err := getFromRequest[T](url)
		if err != nil {
			return res.json, err
		}
		cache.Add(url, res.raw)
		return res.json, nil
	}
	var res T
	json.Unmarshal(cached, &res)
	return res, nil
}

func getFromRequest[T any](url string) (JsonRes[T], error) {
	resp, err := http.Get(url)
	if resp.StatusCode > 399 {
		errorMessage := fmt.Sprintf("Error fetching %s, status code %d", url, resp.StatusCode)
		return JsonRes[T]{}, errors.New(errorMessage)
	}
	if err != nil {
		return JsonRes[T]{}, err

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return JsonRes[T]{}, err
	}
	var res T

	json.Unmarshal(body, &res)
	return JsonRes[T]{json: res, raw: body}, nil
}
