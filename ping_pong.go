package main

import (
	"fmt"
	"time"
)

func ping(message chan string) {
	for {
		fmt.Println("Received", <-message, "Sending PONG")
		message <- "PONG"
		time.Sleep(300 * time.Millisecond)
	}
}

func pong(message chan string) {
	for {
		fmt.Println("Received", <-message, "Sending PING")
		message <- "PING"
		time.Sleep(300 * time.Millisecond)
	}
}

func manager() {
	message := make(chan string)
	go pong(message)
	go ping(message)
	// message <- "PING"
}

func main() {
	go manager()
	var input string
	fmt.Scanln(&input)
}
