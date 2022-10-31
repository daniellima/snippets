package main

import "fmt"

func mainCompositeLiteral() {
	type Car struct {
		Color string
		Brand string
	}

	type Shirt struct {
		Color string
		Size  string
	}

	car := Car{Color: "Blue", Brand: "Tesla"}
	short := Shirt{Color: "Red", Size: "G"}

	fmt.Println(car)
	fmt.Println(short)
	// type ColoredThing struct {
	// 	Color string
	// }

	// type Car struct {
	// 	ColoredThing
	// 	Brand string
	// }
	// type Shirt struct {
	// 	ColoredThing
	// 	Size string
	// }

	// car := Car{Color: "Blue", Brand: "Tesla"}
	// short := Shirt{Color: "Red", Size: "G"}

	// fmt.Println(car)
	// fmt.Println(short)
	// car := &Car{
	// 	ColoredThing: ColoredThing{
	// 		Color: "Blue",
	// 	},
	// 	Brand: "Tesla",
	// }

	// shirt := &Shirt{
	// 	ColoredThing: ColoredThing{
	// 		Color: "Red",
	// 	},
	// 	Size: "G",
	// }

	// fmt.Println(car)
	// fmt.Println(shirt)
}
