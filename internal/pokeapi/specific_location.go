package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocation(locationArea string) (LocationAreaResp, error){
    fullURL := baseUrl + "/location-area/" + locationArea

    checkCacheVal, ok := c.cache.Get(fullURL)
    if ok{
        locationAreasResp := LocationAreaResp{}
        err := json.Unmarshal(checkCacheVal, &locationAreasResp)
        if err != nil {
            return LocationAreaResp{}, err
        }
        return locationAreasResp, nil
    }

    req, err := http.NewRequest("GET", fullURL, nil)
    if err != nil {
        return LocationAreaResp{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return LocationAreaResp{}, err
    }

    defer res.Body.Close()

    if res.StatusCode > 399 {
        return LocationAreaResp{}, fmt.Errorf("bad status code %v", res.StatusCode)
    }

    dat, err := io.ReadAll(res.Body)

    if err != nil {
        return LocationAreaResp{}, err
    }

    locationAreasResp := LocationAreaResp{}
    err = json.Unmarshal(dat, &locationAreasResp)
    if err != nil {
        return LocationAreaResp{}, err
    }


    c.cache.Add(fullURL, dat)

    return locationAreasResp, nil

}
