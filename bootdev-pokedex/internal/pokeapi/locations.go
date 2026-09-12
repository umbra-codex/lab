package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreaPage struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (c *Client) ListLocations(pageURL *string) (LocationAreaPage, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		area := LocationAreaPage{}
		err := json.Unmarshal(val, &area)
		if err != nil {
			return LocationAreaPage{}, err
		}

		return area, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaPage{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaPage{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaPage{}, err
	}

	area := LocationAreaPage{}
	err = json.Unmarshal(data, &area)
	if err != nil {
		return LocationAreaPage{}, err
	}

	c.cache.Add(url, data)
	return area, nil
}
