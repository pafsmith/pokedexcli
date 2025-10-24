package main

import (
	"fmt"
	"io"
	"net/http"
)
func commandMap() error {
	url := "https://pokeapi.co/api/v2/location-area?limit=20&offset=0"
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", body)	

	return nil

}
