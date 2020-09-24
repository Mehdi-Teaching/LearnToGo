package main

import (
	"fmt"
	"net/http"
)

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

// method with 2 return value
func divMod(x, y int) (int, int) {
	return x / y, x % y
}

// method return value and an error type
func contentOfUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return resp.Header.Get("Content-Type"), err
}

func main() {
	fmt.Println(greeting("Reader"))
	fmt.Println(summation(2, 65))
	fmt.Println(multiply(2, 5))
	fmt.Println(divMod(8, 3))
	fmt.Println(contentOfUrl("http://golang.org"))
}
