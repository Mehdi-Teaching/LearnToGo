package main

import (
	"fmt"
	"time"
)

func main() {
	// select is like switch but for channels
	c1 := make(chan string)
	c2 := make(chan string)

	// first goroutine send ping continuously
	go func() {
		for {
			c1 <- "Ping"
			time.Sleep(time.Second * 1)
		}
	}()

	// second goroutine send pong continuously
	go func() {
		for {
			c2 <- "Pong"
			time.Sleep(time.Second * 2)
		}
	}()

	// here we receive values from channels using select
	go func() {
		for {
			select {
			case ping := <-c1:
				fmt.Println(ping)
			case pong := <-c2:
				fmt.Println(pong)
			}
		}
	}()

	// enter to stop the program
	var input string
	fmt.Scanln(&input)
}
