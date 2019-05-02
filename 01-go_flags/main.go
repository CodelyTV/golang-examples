package main

import (
	"flag"
	"fmt"
)

var beers = map[string]string{
	"01D9X58E7NPXX5MVCR9QN794CH": "Mad Jack Mixer",
	"01D9X5BQ5X48XMMVZ2F2G3R5MS": "Keystone Ice",
	"01D9X5CVS1M9VR5ZD627XDF6ND": "Belgian Moon",
}

func main() {
	beersFlag := flag.Bool("beers", false, "show beers")
	flag.Usage = func() {
		fmt.Println("This is a cli create for CodelyTV Go course")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *beersFlag {
		fmt.Println(beers)
	}
}
