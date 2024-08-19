package main

import (
	"fmt"
	"sort"
	"strings"
)

// Standard Library available at: https://golang.org/pkg/

func main() {
	greeting := "Hello, World!"

	fmt.Println(strings.Contains(greeting, "Hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "World"))
	fmt.Println(strings.Split(greeting, ",")) // Returns a slice of strings
	// Prints type
	fmt.Printf("Type: %T\n", strings.Split(greeting, ","))
	
	// Does not modify the original string
	fmt.Println("Original string value: ", greeting)

	// Sort package
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages) // Changes the original slice

	index := sort.SearchInts(ages, 1) // Returns the index of the value if found
	// If not found, returns the index where the value should be inserted

	fmt.Println(index)

	names := []string{"Daniel", "Jenny", "Steven", "Alice", "John"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Alice"))
}