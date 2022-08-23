// 220821 re-construct repository

package main

import (
	"fmt"
	"go_study/something"
	"strings"
)

func main() {
	fmt.Printf("Hello world! \n")
	something.SayHello()
	//something.sayBye() -> expected to error

	const name = "andrew"
	//name = "good" -> expected to error
	//const name string = "andrew"
	//const name bool = "andrew" -> expected to error
	fmt.Println(name)

	//var name1 = "andrew"
	//var name1 string = "andrew"
	name1 := "andrew"
	name1 = "good"
	fmt.Printf(name1)
	fmt.Printf("\n")

	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("andrew")
	// totalLength, _ := lenAndUpper("andrew") -> ignoring value by underscore(_)
	// fmt.Println(totalLength) -> would go well if ignoring certain return value
	fmt.Println(totalLength, upperName)

	repeatMe("ruben", "andrew", "good", "bad")
}

// 220823
func multiply(a, b int) int {
	//func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}
