package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	pokeapi "github.com/abdol-ahmed/pokedexcli/internal/pokecache"
)

var cache *pokeapi.Cache

func init() {
	// Initialize the cache when your application starts.
	const cacheInterval = 30 * time.Second
	cache = pokeapi.NewCache(cacheInterval)
}

func (c *Client) ListLocations(pageURL *string) (Locations, error) {
	locations := Locations{}
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if data, exist := cache.Get(url); exist {
		if err := json.Unmarshal(data, &locations); err != nil {
			return locations, err
		}
		return locations, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return locations, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locations, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locations, err
	}

	// 3. Store the response in the cache before returning.
	cache.Add(url, data)

	if err := json.Unmarshal(data, &locations); err != nil {
		return locations, err
	}

	return locations, nil
}

func (c *Client) GetLocationByName(locationName string) (Location, error) {
	locationDetails := Location{}
	url := baseURL + "/location-area/" + locationName

	if data, exist := cache.Get(url); exist {
		if err := json.Unmarshal(data, &locationDetails); err != nil {
			return locationDetails, err
		}
		return locationDetails, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return locationDetails, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationDetails, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationDetails, err
	}

	// 3. Store the response in the cache before returning.
	cache.Add(url, data)

	if err := json.Unmarshal(data, &locationDetails); err != nil {
		return locationDetails, err
	}

	return locationDetails, nil
}
