package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(url *string) (LocationAreasResp, error){
    fullURL := baseUrl + "/location-area"
    if url != nil{
        fullURL  = *url
    }

    checkCacheVal, ok := c.cache.Get(fullURL)
    if ok{
        locationAreasResp := LocationAreasResp{}
        err := json.Unmarshal(checkCacheVal, &locationAreasResp)
        if err != nil {
            return LocationAreasResp{}, err
        }
        return locationAreasResp, nil
    }

    req, err := http.NewRequest("GET", fullURL, nil)
    if err != nil {
        return LocationAreasResp{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return LocationAreasResp{}, err
    }

    defer res.Body.Close()

    if res.StatusCode > 399 {
        return LocationAreasResp{}, fmt.Errorf("bad status code %v", res.StatusCode)
    }

    dat, err := io.ReadAll(res.Body)

    if err != nil {
        return LocationAreasResp{}, err
    }

    locationAreasResp := LocationAreasResp{}
    err = json.Unmarshal(dat, &locationAreasResp)
    if err != nil {
        return LocationAreasResp{}, err
    }


    c.cache.Add(fullURL, dat)

    return locationAreasResp, nil

}
