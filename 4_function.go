package main

import "fmt"

/*
	define function using <func> keyword
	return type comes after function params
*/

func greeting(name string) string {
	return "Hello " + name
}

// another example
func summation(num1 int, num2 int) int {
	return num1 + num2
}

// if params type are same you can do this:
func multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Println(greeting("Reader"))
	fmt.Println(summation(2, 65))
	fmt.Println(multiply(2, 5))
}
