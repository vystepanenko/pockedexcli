package pokedexapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageUrl *string) (LoacationAreasResp, error) {
	// Get all locations
	endpoint := "/location-area?offset=0&limit=20"
	fullURL := fmt.Sprintf("%s%s", baseURL, endpoint)

	if pageUrl != nil {
		fullURL = *pageUrl
	}

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Cache hit")
		var locationResp LoacationAreasResp
		err := json.Unmarshal(data, &locationResp)
		if err != nil {
			return LoacationAreasResp{}, err
		}

		return locationResp, nil
	}

	fmt.Println("Cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LoacationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LoacationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LoacationAreasResp{}, fmt.Errorf(
			"error requisting areas. StatusCode: %d",
			resp.StatusCode,
		)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LoacationAreasResp{}, err
	}

	var locationResp LoacationAreasResp
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LoacationAreasResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationResp, nil
}

func (c *Client) GetArea(areaName string) (LocationArea, error) {
	// Get all locations
	endpoint := "/location-area/" + areaName

	fullURL := fmt.Sprintf("%s%s", baseURL, endpoint)

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Cache hit")
		var areaResp LocationArea
		err := json.Unmarshal(data, &areaResp)
		if err != nil {
			return LocationArea{}, err
		}

		return areaResp, nil
	}

	fmt.Println("Cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf(
			"error requisting area. StatusCode: %d",
			resp.StatusCode,
		)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	var locationResp LocationArea
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	return locationResp, nil
}
