package pokeapi

import(
	"net/http"
	"io"
	"encoding/json"
)


func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	URL := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(URL); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err 
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(body, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemonResp, nil
}