package main

import (
	"fmt"
	"sync"
	"time"
)

func sampleGoRoutine(customWaitGroup *sync.WaitGroup) {
	for counter := 0; counter < 5; counter++ {
		fmt.Println(counter)
	}
	// When this goroutine is finished, mark it as finished.
	defer customWaitGroup.Done()
}

func main() {
	var customWaitGroup sync.WaitGroup

	customWaitGroup.Add(1)
	go sampleGoRoutine(&customWaitGroup)

	// The main Goroutine waits for all Goroutines to finish
	customWaitGroup.Wait()
	for mainCounter := 0; mainCounter < 5; mainCounter++ {
		fmt.Println("main function is printing!")
		time.Sleep(time.Millisecond * 150)
	}
}
