package main

import "fmt"

func main() {
	x := 5
	// using ampersand to store pointer of a value
	y := &x

	println(x, y)
	fmt.Printf("%T\n", y)

	// reading value using pointer
	// use asterisk to read value with pointer
	println(*y)

	// change value using pointer
	*y = 10
	println(x)

}
