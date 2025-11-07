package pokeapi

import (
	"encoding/json"
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

func NewClient(timeout time.Duration) PokeapiClient {
	return PokeapiClient{
		httpClient: http.Client{Timeout: timeout},
		cache:      pokecache.NewCache(),
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
