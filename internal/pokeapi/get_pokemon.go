package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}

	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	respPokemon := RespPokemon{}

	err = json.Unmarshal(dat, &respPokemon)
	if err != nil {
		return RespPokemon{}, err
	}
	return respPokemon, nil
}
