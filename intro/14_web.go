package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	// you can put html in here
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func numbers(w http.ResponseWriter, r *http.Request) {
	// you can put html in here
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(w, "<h1>"+strconv.Itoa(i)+"</h1>")
	}
}

func main() {
	// pattern is route, you can setup multiple routes
	http.HandleFunc("/", index)
	// localhost:3000/numbers
	http.HandleFunc("/numbers", numbers)
	fmt.Println("Starting Server...")
	// go to browser and type localhost:3000
	http.ListenAndServe(":3000", nil)
}
