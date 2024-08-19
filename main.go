package main

import "fmt"

func main() {

	// Array declaration
	var ages [3]int = [3]int{20, 25, 30}
	fmt.Println(ages, len(ages))

	ages[2]	= 35

	names := [4]string{"John", "Doe", "Jane", "Doe"}

	// Slices
	var scores = []int{100, 50, 60}
	scores = append(scores, 85)

	fmt.Println(ages, names, scores)

	// Slice ranges
	// Get range of elements from an array or slice & create a new slice
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Smith")

	fmt.Println(rangeOne)
}