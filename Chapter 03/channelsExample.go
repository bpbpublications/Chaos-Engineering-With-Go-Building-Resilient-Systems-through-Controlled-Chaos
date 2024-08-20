package main

import (
	"fmt"
	"time"
)

func sendDataPackets(messages chan<- string) {
	for counter := 0; counter < 5; counter++ {
		fmt.Println("Sending data: Data", counter+1)
		messages <- fmt.Sprintf("Data %d", counter+1)
		time.Sleep(time.Millisecond * 400)
	}
	close(messages)
}

func receiveDataPackets(messages <-chan string) {
	for message := range messages {
		fmt.Println("Received data:", message)
	}
}

func main() {
	dataPackets := make(chan string)

	go sendDataPackets(dataPackets)
	receiveDataPackets(dataPackets)
}
