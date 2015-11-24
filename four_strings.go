package main

import "fmt"

func inspector(name string, message chan string) {
	val := <-message
	fmt.Printf("%s: %s\n", name, val)
	message <- val
	// fmt.Println("Shutting down", name)
}

func main() {
	messages := make(chan string, 4)

	for i := 0; i < 20; i++ {
		name := fmt.Sprintf("GR %d", i+1)
		go inspector(name, messages)
	}
	messages <- "A"
	messages <- "B"
	messages <- "C"
	messages <- "D"

	var input string
	fmt.Scanln(&input)
}
