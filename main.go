package main

import (
	"fmt"
)

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("Value of x is:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of i is:", i)
	// }

	names := []string{"John", "Doe", "Jane", "Doe"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("The value at index %v is %v \n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("The value at is %v \n", value)
		value = "new value" // Does not change the value in the slice
		// Value is local copy of the value in the slice
	}

	fmt.Println(names)
}