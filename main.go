package main

import "fmt"

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	name := "tifa"

	// updateName(name)

	// fmt.Println("Memory Address of name is: ", &name)

	m := &name
	fmt.Println("Memory Address: ", m)
	fmt.Println("Value at Memory Address: ", *m)

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

}