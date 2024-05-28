package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/adspampin/pokedexcli/internal/pokeapi"
)

type config struct{
    pokeapiClient pokeapi.Client
    nextLocationURL *string
    prevLocationURL *string
}

func startRepl(cfg *config){
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()

        words := cleanInput(scanner.Text())
        if len(words) ==  0{
            continue
        }

        command := words[0]
        newCommand, exists := getCommands()[command]
        if exists {
            err := newCommand.callback(cfg)
            if err !=  nil {
                fmt.Println(err)
            }
            continue
        }else {
            fmt.Println("Unkown Command")
            continue
        }
    }
}


func cleanInput(text string) []string {
    out := strings.ToLower(text)
    words := strings.Fields(out)
    return words
}

type cliCommand struct{
    name string
    description string
    callback func(cfg *config) error
}


func getCommands() map[string]cliCommand{
    return map[string]cliCommand{
        "help":{
            name: "help",
            description: "Display a help message",
            callback: commandHelp,
        },
        "exit":{
            name: "exit",
            description: "Exit the pokedex",
            callback: commandExit,
        },
        "map":{
            name: "map",
            description: "Get the next page of locations",
            callback: commandMap,
        },
        "mapb":{
            name: "mapb",
            description: "Get the previous page of locations",
            callback: commandMapB,
        },
    }


}
