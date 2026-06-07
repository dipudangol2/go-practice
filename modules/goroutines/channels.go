package main

import (
	"fmt"
	"sync"
	"time"
)

func processNum(numChan chan int, wg *sync.WaitGroup, workerId int) {
	defer wg.Done()

	// The range closes once numChan is closed.
	for num := range numChan {
		fmt.Println("Processing Number", num, "Worker ID: ", workerId)
		// Add a slight delay because if there is no delay there is a chance only one worker can process all the channels due to the worker only needing to take the value from the channel.
		time.Sleep(time.Second)
	}
}

func syncDemo(boolChan chan bool) {
	defer func() {
		boolChan <- true
	}()
	fmt.Println("The task has been synchronized and status has been sent to the channel")
	time.Sleep(time.Millisecond * 500)
}

func UnbufferedChannelDemo() {
	// These are the examples for Unbuffered channels. The problem with them is that they are blocking. They wait for the result before executing further.
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
	for num := range 10 {
		numChan <- num
	}

	// It is important to close the channel once the data has been sent completely to avoid deadlock while waiting for the numbers in the channel.
	close(numChan)

	wg.Wait()

	var boolWg sync.WaitGroup
	boolChan := make(chan bool)

	// wg.Go function can be used to not have to call the add and wait functions only wait needs to be called
	boolWg.Go(
		func() {
			syncDemo(boolChan)
		},
	)
	boolWg.Wait()
	fmt.Println(<-boolChan)
}
