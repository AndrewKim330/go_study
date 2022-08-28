// 220821 re-construct repository

package main

import (
	"fmt"
	"go_study/accounts"
	"go_study/something"
	"strings"
)

// 220827
type person struct {
	name    string
	age     int
	favFood []string
}

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

	totalLength2, upperName2 := lenAndUpper2("andrew")
	fmt.Println(totalLength2, upperName2)

	totalLength3, upperName3 := lenAndUpper3("andrew")
	fmt.Println(totalLength3, upperName3)

	// 220825
	total := superAdd(1, 2, 3, 4, 5)
	fmt.Println(total)

	fmt.Println(canIDrink(16))

	// 220826
	fmt.Println(canIDrink2(16))

	a := 2
	b := &a
	fmt.Println(&a, b)
	fmt.Println(*b)
	*b = 20
	fmt.Println(a)

	// 220827
	names := [5]string{"andrew", "ruben", "kim"}
	names[3] = "a"
	//names[4] = "b"
	//names[5] = "c" -> expected to error
	fmt.Println(names)
	names2 := []string{"a", "b", "c"}
	fmt.Println(names2)
	names2 = append(names2, "aaa")
	fmt.Println(names2)

	andrew := map[string]string{"name": "andrew", "age": "35"}
	fmt.Println(andrew)

	for key, value := range andrew {
		fmt.Println(key, value)
	}

	favFood := []string{"raisin"}
	//andrew2 := person{"andrew", 35, favFood} // not fancy..
	andrew2 := person{name: "andrew", age: 35, favFood: favFood} // explicitly
	fmt.Println(andrew2)

	// 220828
	// account := accounts.BankAccount{Owner: "andrew", Balance: 1000} -> possibility of change
	// fmt.Println(account)

	account := accounts.NewAccount("andrew")
	fmt.Println(account)
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

func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func lenAndUpper3(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// 220825
func superAdd(numbers ...int) int {
	//for idx, number := range numbers {
	//	fmt.Println(idx, number)
	//}
	//for i := 0; i < len(numbers); i++ {
	//	fmt.Println(numbers[i])
	//}
	total := 0
	for _, number := range numbers { // ignoring index
		total += number
	}

	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

// 220826
func canIDrink2(age int) bool {
	//switch age {
	//case 10:
	//	return false
	//case 18:
	//	return true
	//}
	//return false

	//switch {
	//case age < 10:
	//	return false
	//case age == 18:
	//	return true
	//}

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false

}
