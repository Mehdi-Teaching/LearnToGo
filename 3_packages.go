package main

/*
	use parenthesis for multi package
	put each package in separate line
	==================
	READ THIS PART WHEN YOU REACHED LINE 26
	SKIP TO LINE 15
	==================
	Add your packages to src folder in Go workstation
	and use them like other packages
	Read more -> https://golang.org/doc/code.html#Organization
*/

import (
	"Playground/strutil"
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Ceil(12.6))

	// from my own package
	fmt.Println(strutil.StrLen("A Very Long Name"))
}
