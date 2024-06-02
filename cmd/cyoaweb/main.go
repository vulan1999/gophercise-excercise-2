package main

import (
	"adventure"
	"flag"
	"fmt"
	"os"
)

func main() {
	file := flag.String("json", "gopher.json", "The json input adventure file")
	flag.Parse()
	fmt.Printf("Using the story in %s. \n", *file)
	// Open the desired file
	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}
	// Decode the json file
	story, err := adventure.JsonStory(f)
	if err != nil {
		panic(err)
	}
	// Parsing the story
	fmt.Printf("%+v\n", story)
}
