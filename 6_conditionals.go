package main

import "fmt"

func main() {
	x := 50
	y := 10

	// if else
	if x > y {
		fmt.Println("X is bigger than Y")
	} else {
		fmt.Println("Y is bigger than X")
	}

	// switch
	color := "red"
	switch color {
	case "red":
		println("Color is Red")
	case "blue":
		println("Color is Blue")
	default:
		println("Color is unknown")
	}
}
