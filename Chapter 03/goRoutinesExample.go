package main

import (
	"fmt"
	"time"
)

func sampleGoRoutine() {
	for counter := 0; counter < 5; counter++ {
		fmt.Println("sampleGoRoutine function is printing!")
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// This initiates a Goroutine that runs in parallel
	// with the primary Goroutine(which is main Goroutine)
	go sampleGoRoutine()

	// The main Goroutine continues executing.
	for mainCounter := 0; mainCounter < 5; mainCounter++ {
		fmt.Println("main function is printing!")
		time.Sleep(time.Millisecond * 150)
	}
}
