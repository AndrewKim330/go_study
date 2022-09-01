package main

import (
	"fmt"
	"go_study/dict"
)

func main() {
	//dictionary := dict.Dictionary{"first": "First word"}
	//// fmt.Println(dictionary["first"]) -> not fancy....
	//definition, err := dictionary.Search("first")
	//// definition, err := dictionary.Search("second") // expect to print error
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(definition)
	//}

	// 220831
	word := "hello"
	definition := "greeting"
	dictionary := dict.Dictionary{}
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, err := dictionary.Search(word)
	fmt.Println("found: ", word, "definition: ", hello)

	err2 := dictionary.Add(word, definition)
	fmt.Println(err2) // -> expected to error text

	err3 := dictionary.Update(word, "Second")
	if err3 != nil {
		fmt.Println(err3)
	}
	ww, _ := dictionary.Search(word)
	fmt.Println(ww)

	err4 := dictionary.Update("wwwwwww", "Third")
	if err4 != nil {
		fmt.Println(err4)
	}
	www, _ := dictionary.Search(word)
	fmt.Println(www)

	dictionary.Delete(word)
	wwww, err5 := dictionary.Search(word)
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Println(wwww)

}
