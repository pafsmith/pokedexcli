package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) FindPokemon(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name

	if name == "" {
		return RespPokemon{}, errors.New("pokemon can not be blank")
	}

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespPokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespPokemon{}, err
		}
		return pokemonResp, nil
	}
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

	pokemonResp := RespPokemon{}

	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
