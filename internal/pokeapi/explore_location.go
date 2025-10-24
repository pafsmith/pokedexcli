package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)


func (c *Client) ExploreLocation(location string) (RespExploreLocation, error) {

	if location == "" {
		return RespExploreLocation{}, errors.New("no argument provided")

	}
	url := baseURL + "/location-area/" + location 

	if val, ok := c.cache.Get(url); ok {
		exploreLocationResp := RespExploreLocation{}
		err := json.Unmarshal(val, &exploreLocationResp)
		if err != nil {
			return RespExploreLocation{}, err
		}
		return exploreLocationResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespExploreLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespExploreLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExploreLocation{}, err
	}

	exploreLocationResp := RespExploreLocation{}

	err = json.Unmarshal(dat, &exploreLocationResp)
	if err != nil {
		return RespExploreLocation{}, err
	}

	c.cache.Add(url, dat)
	return exploreLocationResp, nil
} 
