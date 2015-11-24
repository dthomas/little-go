package main

import (
	"fmt"
	"time"
)

func main() {
	baton := make(chan int)
	go runner(baton)
	baton <- 1

	// fmt.Println("Press any key to continue...")
	var input string
	fmt.Scanln(&input)
}

func runner(baton chan int) {
	running := <-baton
	var newRunner int

	fmt.Printf("Runner %d running with baton\n", running)
	time.Sleep(1000 * time.Millisecond)

	if running != 4 {
		newRunner = running + 1
		go runner(baton)
		fmt.Printf("Runner %d reaching line\n", newRunner)
	} else if running == 4 {
		fmt.Println("Finishing Race")
		return
	}

	fmt.Printf("Runner %d Exchanging baton with %d\n", running, newRunner)
	baton <- newRunner
}
