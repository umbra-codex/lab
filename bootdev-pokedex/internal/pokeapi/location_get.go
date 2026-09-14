package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(locationName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		area := LocationArea{}
		err := json.Unmarshal(val, &area)
		if err != nil {
			return LocationArea{}, err
		}

		return area, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("%s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	area := LocationArea{}
	err = json.Unmarshal(data, &area)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, data)
	return area, nil
}
