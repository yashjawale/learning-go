package main

import "fmt"

func main() {
	age := 35
	name := "yash"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hello world!")
	fmt.Println("goodbye ninjas!")
	fmt.Println("my age is", age, "and name is", name)

	// Printf (formatted strings) %_ = format specifier
	fmt.Printf(("my age is %v and name is %v \n"), age, name)
	fmt.Printf(("my age is %q and name is %q \n"), age, name)
	fmt.Printf(("age is of type %T \n"), age)
	fmt.Printf(("you scored %f points! \n"), 9.02)
	fmt.Printf(("you scored %0.1f points! \n"), 9.0)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf(("my age is %v and name is %v \n"), age, name)
	fmt.Println("the saved string is:", str)
}