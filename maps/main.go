package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	addColor(&colors, "black", "#000000")
	addColor(&colors, "blue", "#0000ff")
	printMap(colors)
}

func addColor(c *map[string]string, color string, hex string) {
	(*c)[color] = hex
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
