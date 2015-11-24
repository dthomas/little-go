package main

import "fmt"

/*
// Write a program where two goroutines pass an integer back and forth ten times.
// Display when each goroutine receives the integer. Increment the integer with
// each pass. Once the integer equals ten, terminate the program cleanly.
*/

func tosser(name string, num chan int) {
	for {
		val, ok := <-num
		if !ok {
			fmt.Println("Closing", name)
			return
		}

		fmt.Printf("%s: %d\n", name, val)

		if val == 10 {
			fmt.Println("Closing", name)
			close(num)
			return
		}

		num <- (val + 1)
	}
}

func main() {
	num := make(chan int)
	go tosser("Luv", num)
	go tosser("Kush", num)
	num <- 1

	var input string
	fmt.Scanln(&input)
}
