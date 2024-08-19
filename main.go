package main

import "fmt"

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	// Non pointer values

	name := "John"
	name = updateName(name)
	fmt.Println(name)

	// group B types -> slices, maps, functions
	// Pointer wrapper values
	menu := map[string]float64{
		"soup":  4.99,
		"pie":   7.99,
		"salad": 6.99,
		"coffee": 1.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}