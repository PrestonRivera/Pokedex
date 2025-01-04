package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
	"github.com/PrestonRivera/Pokedex/internal/pokecache"
)


const baseURL = "https://pokeapi.co/api/v2"


func (c *Client)ListLocationAreas(pageURL *string) (respLocations, error) {
	URL := baseURL + "/location-area"
	if pageURL != nil {
		URL = *pageURL
	}
	
	cache := pokecache.NewCache(time.Minute * 5)
	if data, found := cache.Get(URL); found {
		var locations respLocations
		if err := json.Unmarshal(data, &locations); err != nil {
			return respLocations{}, err
		}
		return locations, nil
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return respLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return respLocations{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return respLocations{}, err
	}

	var locations respLocations
	if err := json.Unmarshal(data, &locations); err != nil {
		return respLocations{}, err
	}
	return locations, nil
} 