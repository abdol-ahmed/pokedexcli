package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	pokeapi "github.com/abdol-ahmed/pokedexcli/internal/pokecache"
)

func init() {
	// Initialize the cache when your application starts.
	const cacheInterval = 30 * time.Second
	cache = pokeapi.NewCache(cacheInterval)
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	pokemon := Pokemon{}
	url := baseURL + "/pokemon/" + pokemonName

	if data, exist := cache.Get(url); exist {
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return pokemon, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return pokemon, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return pokemon, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemon, err
	}

	// 3. Store the response in the cache before returning.
	cache.Add(url, data)

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return pokemon, err
	}

	return pokemon, nil
}
