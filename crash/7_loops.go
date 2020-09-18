package main

import "fmt"

func main() {

	// long method for loop
	i := 1
	for i <= 10 {
		println(i)
		i++
	}

	// short way for loop
	for i := 1; i <= 10; i++ {
		fmt.Printf("i is %d\n", i)
	}

	println("=========")

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			println("FizzBuzz")
		} else if i%3 == 0 {
			println("Fizz")
		} else if i%5 == 0 {
			println("Buzz")
		} else {
			println(i)
		}
	}
}
