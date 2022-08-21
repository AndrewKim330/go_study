// 220821 re-construct repository

package main

import (
	"fmt"
	"go_study/something"
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

}
