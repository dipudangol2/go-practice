package main

import (
	"fmt"
	"sync"
)

func processNum(numChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Processing Number", <-numChan)
}

func channelDemo() {
	fmt.Println("Channels in go")

	// Create a channel with make specifying the type of data sent through the channel
	// This creates a deadlock
	// messageChannel := make(chan string)
	//
	// messageChannel <- "ping"
	//
	// msg := <-messageChannel
	//
	// fmt.Println(msg)
	var wg sync.WaitGroup
	numChan := make(chan int)

	wg.Add(1)
	go processNum(numChan, &wg)
	numChan <- 5
	wg.Wait()
}
