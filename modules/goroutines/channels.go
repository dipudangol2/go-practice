package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processNum(numChan chan int, wg *sync.WaitGroup, workerId int) {
	defer wg.Done()

	// The range closes once numChan is closed.
	for num := range numChan {
		fmt.Println("Processing Number", num, "Worker ID: ", workerId)
		// Add a slight delay because if there is no delay there is a chance only one worker can process all the channels due to the worker only needing to take the value from the channel.
		time.Sleep(time.Millisecond * 40)
	}
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

	// Initialize a waitgroup and channel for integer
	var wg sync.WaitGroup
	numChan := make(chan int)

	// We want to have multiple workers for the processing of data in the channel
	// We Initialize the number of workers and start the execution for all the workers
	workerNum := 5
	for i := 1; i <= workerNum; i++ {
		wg.Add(1)
		go processNum(numChan, &wg, i)
	}

	// send 30 items to the pipeline
	for range 30 {
		numChan <- rand.Intn(100)
	}

	// It is important to close the channel once the data has been sent completely to avoid deadlock while waiting for the numbers in the channel.
	close(numChan)

	wg.Wait()
}
