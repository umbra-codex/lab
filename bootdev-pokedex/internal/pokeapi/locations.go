package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		list := LocationAreaList{}
		err := json.Unmarshal(val, &list)
		if err != nil {
			return LocationAreaList{}, err
		}

		return list, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaList{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaList{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaList{}, err
	}

	list := LocationAreaList{}
	err = json.Unmarshal(data, &list)
	if err != nil {
		return LocationAreaList{}, err
	}

	c.cache.Add(url, data)
	return list, nil
}
