package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *config) error {

	res, err := http.Get(cfg.Next)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &cfg)
	if err != nil {
		return err
	}

	locations := cfg.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	res, err := http.Get(cfg.Previous.(string))
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &cfg)
	if err != nil {
		return err
	}

	locations := cfg.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}
