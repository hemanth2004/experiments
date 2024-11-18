package main

import "fmt"

// nested structs
type car struct {
	brand      string
	model      string
	doors      int
	frontWheel wheel
	backWheel  wheel
}

type wheel struct {
	radius float64
}

// anonymous structs
func test() {
	weapon := struct {
		make  string
		model string
		wheel int
	}{
		make:  "tesla",
		model: "3",
		wheel: 4,
	}

	fmt.Println(weapon)
}

// embedded structs
type sedan struct {
	car
	trunkModel string
	sedanClass string
}

// size of empty structs = 0 bytes
type empt struct{}

func return0bytes() empt {
	return empt{}
}

func main() {
	myCar := car{"tata", "nexus", 4, wheel{}, wheel{}}
	myCar.frontWheel.radius = 64.0
	//fmt.Println(myCar.brand)

	mySedan := sedan{
		car: car{
			brand:      "tata",
			model:      "nexus",
			doors:      4,
			frontWheel: wheel{},
			backWheel:  wheel{radius: 64.0},
		},
		trunkModel: "tata",
		sedanClass: "S",
	}

	sizeof(mySedan)
	fmt.Println(mySedan)
}
