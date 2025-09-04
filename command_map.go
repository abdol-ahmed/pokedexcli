package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationArea struct {
	Name string
	Url  string
}

type config struct {
	Count    int
	Next     string
	Previous string
	Results  []locationArea
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != "" {
		url = cfg.Next
	}

	return getLocationAreas(url, cfg)
}

func commandMapb(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Previous != "" {
		url = cfg.Previous
	}

	return getLocationAreas(url, cfg)
}

func getLocationAreas(url string, cfg *config) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return err
	}

	for _, area := range cfg.Results {
		fmt.Println(area.Name)
	}
	return nil
}
