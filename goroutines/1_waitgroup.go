package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// print url and url type
	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://w3schools.com",
	}

	// doing normally
	fmt.Println("doing normally")
	for _, url := range urls {
		printUrlType(url)
	}

	// using goroutines
	println("using goroutines")

	// use WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup
	for _, url := range urls {
		// add to waitGroup for each goroutine
		wg.Add(1)
		// create a goroutine
		go func(url string) {
			printUrlType(url)
			// after the work done
			wg.Done()
		}(url)
	}

	// wait for all goroutines to finish
	wg.Wait()

}

func printUrlType(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("%s !!! Nothing found \n", url)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s \n", url, ctype)

}
