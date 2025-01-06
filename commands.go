package main

import (
	"fmt"
	"os"
	"errors"
	"math/rand"
	"maps"
	"github.com/PrestonRivera/Pokedex/internal/pokeapi"
)


type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}


type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
	pokeapiClient pokeapi.Client
	caughtPokemon map[string]pokeapi.Pokemon
}


func commandHelp(cfg *config, args ...string) error {
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


func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandMap(cfg *config, args ...string) error {
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


func commandMapb(cfg *config, args ...string) error {
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


func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a location name")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.GetLocations(name)
	if err != nil {
		return err 
	}
	fmt.Printf("Exploring %s...", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}


func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a Pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	userChance := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s....\n", pokemon.Name)
	if userChance > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}


func commandInspect(cfg *config, args ...string) error {
	 name := args[0]

	 value, key := cfg.caughtPokemon[name]
	 if key {
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", value.Name, value.Height, value.Weight)
		fmt.Println("Stats:")
		for _, stat := range value.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range value.Types {
			fmt.Printf("  -%s\n", t.Type.Name)
		}
	 } else {
		fmt.Println("You have not caught this Pokemon")
	 }
	 return nil
}


func commandPokedex(cfg *config, args ...string) error {
	keys := maps.Keys(cfg.caughtPokemon)

	fmt.Println("Your Pokedex:")
	for k := range keys {
		fmt.Printf(" - %s\n", k)
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
		"explore":{
			name: "explore <location_name>",
			description: "Explore a location",
			callback: commandExplore,
		},
		"catch":{
			name: "catch <pokemon_name>",
			description: "Used to catch Pokemon",
			callback: commandCatch,
		},
		"inspect":{
			name: "inspect <pokemon_name>",
			description: "Used to see details about Pokemon you have caught",
			callback: commandInspect,
		},
		"pokedex":{
			name: "pokedex",
			description: "Lists all the Pokemon you have caught",
			callback: commandPokedex,
		},
	}
}