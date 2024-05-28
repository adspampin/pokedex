package main

import (
	"errors"
	"fmt"
)


func commandMap(cfg *config, args ...string) error {
    locationResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
    if err != nil {
        return err
    }

    cfg.nextLocationURL = locationResp.Next
    cfg.prevLocationURL = locationResp.Previous

    for _, loc := range locationResp.Results {
        fmt.Println(loc.Name)
    }

    return nil
}


func commandMapB(cfg *config, args ...string) error {

    if cfg.prevLocationURL == nil {
        return errors.New("you're on the first page")
    }

    locationResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationURL)
    if err != nil {
        return err
    }

    cfg.nextLocationURL = locationResp.Next
    cfg.prevLocationURL = locationResp.Previous

    for _, loc := range locationResp.Results {
        fmt.Println(loc.Name)
    }

    return nil

}
