package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)


func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		if len(input) < 0 {
			continue
		}
		commandName := input[0]
		
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println("err")
			}
			continue
		} else {
			fmt.Println("command unknown")
			continue
		}
	}
}


func cleanInput(text string) []string {
	input := strings.ToLower(text)
	output :=  strings.Fields(input)
	return output
}