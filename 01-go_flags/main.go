package main

import (
	"fmt"
	"os"
)

var beers = map[string]string{
	"01D9X58E7NPXX5MVCR9QN794CH": "Mad Jack Mixer",
	"01D9X5BQ5X48XMMVZ2F2G3R5MS": "Keystone Ice",
	"01D9X5CVS1M9VR5ZD627XDF6ND": "Belgian Moon",
}

func main() {
	param := os.Args[1]
	if param == "beers" {
		fmt.Println(beers)
	}
	fmt.Println(param)
}
