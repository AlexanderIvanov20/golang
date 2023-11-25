package main

import (
	"fmt"
)

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("The %s color has the %s hex code\n", color, hex)
	}
}

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#ffffff",
	}

	printMap(colors)
}
