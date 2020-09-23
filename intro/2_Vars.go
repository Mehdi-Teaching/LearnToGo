package main

import "fmt"

func main() {
	// Main data types
	/*
		- string
		- bool
		- int
			- int int8 int16 int32 int64
			- uint uint8 uint16 uint32 uint64
		- byte (alias for uint8)
		- rune (alias for uint32)
		- float
			- float32 float64
		- complex
			- complex64 complex128
	*/

	// define variable using var
	// variable type is not necessary
	var name string = "Mehdi"
	var age int = 22
	var isMale = true

	// printing variables
	fmt.Println(name, age, isMale)

	// print variable type
	fmt.Printf("%T\n", age)
	// refer to https://golang.org/pkg/fmt/ for format list

	// creating constant variable
	const phoneNumber = 989121234567
	fmt.Println(phoneNumber)

	// second way to create a variable
	// only works in functions
	height := 7.6
	fmt.Println(height)
	fmt.Printf("%T\n", height)

	// multi assignment
	firstName, lastName := "Mehdi", "Teymorian"
	println(firstName, lastName)

}
