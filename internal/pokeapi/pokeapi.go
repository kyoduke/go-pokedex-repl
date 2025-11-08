package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/kyoduke/pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

type PokeapiClient struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) PokeapiClient {
	return PokeapiClient{
		httpClient: http.Client{Timeout: timeout},
		cache:      pokecache.NewCache(cacheInterval),
	}
}

func (c *PokeapiClient) ListLocationAreas(pageURL *string) (RespLocation, error) {
	fullURL := baseURL + "/location-area?offset=0"
	if pageURL != nil {
		fullURL = *pageURL
	}

	if val, ok := c.cache.Get(fullURL); ok {
		var respLocations RespLocation

		if err := json.Unmarshal(val, &respLocations); err != nil {
			return RespLocation{}, err
		}

		return respLocations, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return RespLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer res.Body.Close()

	var respLocations RespLocation

	respData, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocation{}, err
	}

	c.cache.Add(fullURL, respData)

	if err := json.Unmarshal(respData, &respLocations); err != nil {
		return RespLocation{}, err
	}

	return respLocations, nil
}

func (c *PokeapiClient) ListAreaEncounters(area *string) (RespAreaInfo, error) {
	if area == nil {
		return RespAreaInfo{}, errors.New("Error when calling ListAreaEncounters: area cannot be nil")
	}

	fullURL := baseURL + "location-area/" + *area

	if val, ok := c.cache.Get(fullURL); ok {
		var respAreaInfo RespAreaInfo
		json.Unmarshal(val, &respAreaInfo)

		return respAreaInfo, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return RespAreaInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespAreaInfo{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return RespAreaInfo{}, fmt.Errorf("%s area not found", *area)
	}

	respData, err := io.ReadAll(res.Body)
	if err != nil {
		return RespAreaInfo{}, err
	}

	c.cache.Add(fullURL, respData)

	var respAreaInfo RespAreaInfo
	json.Unmarshal(respData, &respAreaInfo)

	return respAreaInfo, nil
}

func (c *PokeapiClient) CatchPokemon(pokemonName string) (RespPokemon, error) {
	fullURL := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(fullURL); ok {
		var respPokemon RespPokemon
		err := json.Unmarshal(val, &respPokemon)
		if err != nil {
			return RespPokemon{}, err
		}

		return respPokemon, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return RespPokemon{}, fmt.Errorf("pokemon with name '%s' could not be found", pokemonName)
	}

	respData, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemon{}, err
	}
	c.cache.Add(fullURL, respData)

	var respPokemon RespPokemon
	err = json.Unmarshal(respData, &respPokemon)
	if err != nil {
		return RespPokemon{}, err
	}

	return respPokemon, nil
}
