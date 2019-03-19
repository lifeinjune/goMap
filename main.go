package main

import "fmt"

func main() {
	colors := map[string]string{ //declaring map
		"red":   "#ff0000", //assign key and value pair in map
		"green": "#br6551", //assign key and value pair in map
	}
	//var color map[string]string    //another way to make map with empty value
	col := make(map[string]string) //another way to make map with empty value
	col["white"] = "#fffff"

	delete(colors, "red")
	fmt.Println(col)
	fmt.Println(colors)
}
