package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":  4.99,
		"salad": 6.99,
		"pasta": 8.99,
		"pizza": 9.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["salad"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		1234567890: "John",
		9987654321: "Jane",
		5894544530: "Joe",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[1234567890])

	phonebook[1234567890] = "John Doe"
	fmt.Println(phonebook)

	phonebook[5894544530] = "Yoshi"
	fmt.Println(phonebook)
}