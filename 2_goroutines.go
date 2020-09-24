package main

import (
	"fmt"
	"net/http"
)

func main() {
	// doing previous example without synchronization and only using channels

	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://w3schools.com",
	}

	// goroutines communicate through channels
	// create a channel using chan keyword
	urlChannel := make(chan string)
	// using channel
	for _, url := range urls {
		// inside a goroutine, send values to channel
		// or in another word channel send values
		go func(url string) { urlChannel <- urlType(url) }(url)
	}

	i := 0
	for {
		// receive values from channel
		// channel is blocked until result is ready to receive the value
		result := <-urlChannel
		fmt.Println(result)
		i++
		if i == len(urls) {
			break
		}
	}

	// the above example is bi-directional channel
	// you can restrict the channel
	// (c chan <- string ) in this notation, only sending is possible
	// (c <- chan string) only receiving is possible
}

func urlType(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("%s !!! Nothing found \n", url)
		return ""
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	return fmt.Sprintf("%s -> %s \n", url, ctype)
}

// you can only send value through c
func pongChannel(c chan<- string) {
	go func(c chan<- string) { c <- "Pong" }(c)
}

// you can only receive value from c
func pingChannel(c <-chan string) {
	msg := <-c
	fmt.Println(msg)
}
