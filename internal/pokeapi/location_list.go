package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Client) FetchLocationAreas(url string) (PokeAPI, error) {
	cacheData, ok := h.cache.Get(url)

	if ok {
		var poke PokeAPI

		if err := json.Unmarshal(cacheData, &poke); err != nil {
			return PokeAPI{}, err
		}

		return poke, nil
	}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return PokeAPI{}, err
	}

	res, err := h.httpClient.Do(req)

	if err != nil {
		return PokeAPI{}, err
	}

	defer res.Body.Close()

	var pokeResult PokeAPI

	err = json.NewDecoder(res.Body).Decode(&pokeResult)

	if err != nil {
		return PokeAPI{}, err
	}

	marshalData, _ := json.Marshal(pokeResult)
	h.cache.Add(url, marshalData)

	return pokeResult, nil

}

func (h *Client) FetchExplore(name string) (LocationArea, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)

	fmt.Println("Recive name " + name)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return LocationArea{}, nil
	}

	res, err := h.httpClient.Do(req)

	if err != nil {
		return LocationArea{}, nil
	}
	defer res.Body.Close()

	var pokeResult LocationArea

	err = json.NewDecoder(res.Body).Decode(&pokeResult)

	if err != nil {
		return LocationArea{}, nil
	}

	return pokeResult, nil

}
