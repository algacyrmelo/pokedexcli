package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocationAreas returns a list of location areas from PokeAPI,
// using a cached response if one is available.
func (c *Client) ListLocationAreas(pageURL *string) (NamedAPIResourceList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, cached := c.cache.Get(url); cached {
		locationAreas := NamedAPIResourceList{}
		err := json.Unmarshal(data, &locationAreas)
		if err != nil {
			return NamedAPIResourceList{}, err
		}
		return locationAreas, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return NamedAPIResourceList{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	locationAreas := NamedAPIResourceList{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	c.cache.Add(url, data)
	return locationAreas, nil
}
