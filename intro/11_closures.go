package main

import "fmt"

// return type of adder function is a function

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// closures are anonymous function

	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
