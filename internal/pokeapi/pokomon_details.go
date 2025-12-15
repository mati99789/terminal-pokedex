package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Client) PokemonDetails(name string) (Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", name)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := h.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	var pokomon Pokemon

	err = json.NewDecoder(res.Body).Decode(&pokomon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokomon, nil
}
