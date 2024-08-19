package main

import "fmt"

var score = 99.5

func main() {
	sayHello("Gopher")

	for _, p := range points {
		fmt.Println(p)
	}

	// var score = 99.5 // This will not work
	// Has to be declared outside of the main function at root level

	showScore()
}