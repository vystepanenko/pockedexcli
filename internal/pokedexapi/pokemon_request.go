package pokedexapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	// Get all locations
	endpoint := "/pokemon/" + pokemonName

	fullURL := fmt.Sprintf("%s%s", baseURL, endpoint)

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Cache hit")
		var pokemon Pokemon
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	fmt.Println("Cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf(
			"error requisting pokemon. StatusCode: %d",
			resp.StatusCode,
		)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemon, nil
}
