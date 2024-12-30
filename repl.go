package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)


func removePunctuation(word string) string {
	word = strings.Trim(word, ",")
	return word
}


func cleanInput(text string) []string {
	lowerString := strings.ToLower(text)
	input := strings.Fields(strings.TrimSpace(lowerString))
	cleanSlice := []string{}

	for _, word := range input {
		cleanWord := removePunctuation(word)
		if cleanWord != "" {
			cleanSlice = append(cleanSlice, cleanWord)
		}
	}
	return cleanSlice
}


func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		command := input[0]

		commandName, exists := getCommands()[command] 
		if exists {
			err := commandName.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}


type cliCommand struct {
	name string
	description string
	callback func()error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Display a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}