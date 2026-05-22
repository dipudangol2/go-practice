package main

import (
	"fmt"
	"sync"
)

// Task : Whenever you want to use wait group use pointer for referencing the original waitgroup variable created
func Task(i int, wg *sync.WaitGroup) {
	// The waitgroup is used to wait for the completion of the goroutines. The done method is used to signal that the goroutine is completed and the waitgroup can decrement the counter. The defer keyword is used to ensure that the done method is called after the function is completed, even if there is an error or panic in the function. This way, we can ensure that the waitgroup will not be stuck waiting for a goroutine that has already completed.
	defer wg.Done()
	fmt.Println("Task ", i)
}

// Demo of waitGroups for multithreading with goroutines
func waitGrp() {
	fmt.Println("Go routines using waitgroups for better control")
	// Create waitgroup
	var wg sync.WaitGroup

	for idx := range 10 {
		// Increment the number of counter for the goroutine to wait for. This should be done before the go routine is called, otherwise the main thread might end before the counter is incremented and the waitgroup will not work as expected.
		wg.Add(1)
		// Pass the address of the waitgroup for the pointer parameter of the function
		go Task(idx, &wg)
	}

	// Wait until the goroutines are completed. The wait method will block the main thread until the counter of the waitgroup is zero, which means that all the goroutines have completed their tasks and called the done method to decrement the counter.
	wg.Wait()
}
