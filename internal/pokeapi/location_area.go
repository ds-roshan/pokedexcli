package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(locationName string) (RespLocationArea, error) {
	url := baseURL + "/location-area" + "/" + locationName

	fmt.Printf("Exploring %v...\n", locationName)

	if val, ok := c.cache.Get(url); ok {
		respLocationArea := RespLocationArea{}
		err := json.Unmarshal(val, &respLocationArea)
		if err != nil {
			return RespLocationArea{}, err
		}
		return respLocationArea, nil
	}

	req, error := http.NewRequest("GET", url, nil)
	if error != nil {
		return RespLocationArea{}, error
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, nil
	}

	respLocationArea := RespLocationArea{}

	err = json.Unmarshal(dat, &respLocationArea)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, dat)

	return respLocationArea, nil
}
