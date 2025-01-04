package main

import (
	"fmt"
	"os"

	"github.com/PrestonRivera/Pokedex/internal/pokeapi"
)


type cliCommand struct {
	name string
	description string
	callback func(cfg *config) error
}


type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
	pokeapiClient pokeapi.Client
}


func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("^_^ Welcome to Pokedex ^_^")
	fmt.Println()
	fmt.Println("Command Usage:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %v\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}


func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandMap(cfg *config) error {
	respLocations, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = respLocations.Next
	cfg.prevLocationsURL = respLocations.Previous
	
	for _, loc := range respLocations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}


func commandMapb(cfg *config) error {
	respLocations, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationsURL)
	if err != nil {
		return err 
	}
	cfg.nextLocationsURL = respLocations.Next
	cfg.prevLocationsURL = respLocations.Previous

	for _, loc := range respLocations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help":{
			name: "help",
			description: "Displays a helpful message",
			callback: commandHelp,
		},
		"exit":{
			name: "exit",
			description: "Exits the program",
			callback: commandExit,
		},
		"map":{
			name: "map",
			description: "Displays 20 locations",
			callback: commandMap,
		},
		"mapb":{
			name: "mapb",
			description: "Displays 20 previous locations",
			callback: commandMapb,
		},
	}
}