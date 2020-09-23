package main

import "fmt"

func main() {
	// Arrays
	// fix length
	// have to specify type
	var people [2]string
	people[0] = "Alex"
	people[1] = "Peter"
	fmt.Println(people)

	// declare and assign at the same time

	grades := [3]int{19, 16, 14}
	fmt.Println(grades)

	// access a specific element
	fmt.Println(grades[0])

	// ====================

	// Slices
	// pretty much the same
	// difference : slices doesn't have size
	fruits := []string{"Apple", "Orange", "Golabi"}
	fmt.Println(fruits)
	fmt.Printf("%T\n", fruits)
	// part of slice
	fmt.Println(fruits[0:2])

}
