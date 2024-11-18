package main

import "fmt"

type StringMap = map[string]string

/*
type x = y is different from #define x y
1. scope of variable
2. compilation vs preprocessing
3.
*/

type Age int

func (a Age) isAdult() bool {
	return a >= 18
}

func main() {
	ageToTest := 21
	//myMap := StringMap{"hello": "world"}
	//fmt.Println(myMap) // Output: map[hello:world]

	fmt.Println(Age(ageToTest), "is an adult?", Age(18).isAdult()) // Output: 18
}
~