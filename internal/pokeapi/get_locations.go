package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
)



func (c *Client) GetLocations(locationName string) (ExploredLoc, error) {
	URL := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(URL); ok {
		exploredloc := ExploredLoc{}
		if err := json.Unmarshal(val, &exploredloc); err != nil {
			return ExploredLoc{}, err
		}
		return exploredloc, nil
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return ExploredLoc{}, err 
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ExploredLoc{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploredLoc{}, err
	}

	exploredLoc := ExploredLoc{}
	err = json.Unmarshal(body, &exploredLoc)
	if err != nil {
		return ExploredLoc{}, err
	}

	c.cache.Add(URL, body)

	return exploredLoc, nil
}