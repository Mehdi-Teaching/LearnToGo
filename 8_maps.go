package main

import "fmt"

func main() {
	// Maps are key value pairs
	// Define map
	// make(map[KEY-TYPE]VALUE-TYPE)
	emails := make(map[string]string)
	// assign
	emails["Me"] = "Me@Home.com"
	emails["You"] = "You@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Me"])
	fmt.Println(len(emails))

	// delete from map
	delete(emails, "You")
	fmt.Println(emails)

	// second way to create maps
	people := map[string]string{
		"Mehdi": "Teymorian",
		"Alex":  "Johnas"}

	fmt.Println(people)
}
