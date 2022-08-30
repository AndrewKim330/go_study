package main

import (
	"fmt"
	"go_study/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "First word"}
	// fmt.Println(dictionary["first"]) -> not fancy....
	definition, err := dictionary.Search("first")
	// definition, err := dictionary.Search("second") // expect to print error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
