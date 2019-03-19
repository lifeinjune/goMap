package main

import "fmt"

func main() {
	colors := map[string]string{ //declaring map
		"red":   "#ff0000", //assign key and value pair in map
		"green": "#br6551", //assign key and value pair in map
		"white": "#fffff",
	}
	//var color map[string]string    //another way to make map with empty value
	// col := make(map[string]string) //another way to make map with empty value
	// col["white"] = "#fffff" //assign value to map

	// delete(colors, "red") //delete value from the map
	// fmt.Println(col)
	printMap(colors)
}

func printMap(c map[string]string) {

	for key, value := range c {
		fmt.Println(key, value)
	}
}
