package main

import (
	"fmt"
	"os"
)

func main() {
	//panicCode()

	// =====================
	file := openFile("intro/15_panic.go")
	println(file.Name())

}

func panicCode() {
	vals := []int{1, 2, 3}
	// this line cause a panic: index out of bound
	v := vals[10]
	fmt.Println(v)
}

// this method return a file
// if file not found it panic!
func openFile(addr string) os.File {
	file, err := os.Open(addr)

	if err != nil {
		panic(err)
	}

	defer file.Close() // run after method ends

	fmt.Println(" file opened")
	return *file
}
