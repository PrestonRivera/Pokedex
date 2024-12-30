package main

import (
	"fmt"
)


func commandHelp() error {
	fmt.Println("----------")
	fmt.Println("^_^ Welcome to the Pokedex ^_^")
	fmt.Println("----------")
	fmt.Println("Command Usage: ")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("----------")
	return nil
}