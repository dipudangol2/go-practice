package main

import (
	"fmt"
	"sync"
)

/*
* Mutex
* */

// Example of mutex using a struct post
// Mutex is a synchronization technique which locks a specific resource until the operation on it is completed and then releases it. It follows a strict ownership where the thread which locks the resource is only allowed to unlock the resource for others to access
// While mutex does help in avoiding race condition it can also cause bottlenecks as the resource is held up by the goroutine which accesses it first. We can use channels for better performance on larger number of goroutines and complex operations. Mutex are preferred on atomic and smaller operations
type post struct {
	// Create a mutex. After a mutex is declared inside a struct the struct object should always be passed by reference not value or copy.
	mut   sync.Mutex
	views int
}

// create an constructor function for the post struct
func createPost() post {
	p := post{views: 0}
	return p
}

// function which increments the post view count by 1
func (p *post) viewPost() {
	// Unlock the resource after the operation on it is complete. If not done can cause DEADLOCKS
	// defer so that the mutex is unlocked even if the function crashes
	defer p.mut.Unlock()
	// only keep the modification inside the lock so that it doesnot become a burden on speed. for eg if there are other operations keep it above these lines
	// Lock the resource before operating on it
	p.mut.Lock()
	p.views += 1
}

// function which returns Post count
func (p *post) postCount() int {
	return p.views
}

func main() {
	fmt.Println("Mutexes in Go:")

	for run := 0; run < 10; run++ {

		// Create a new post
		post1 := createPost()

		post1.viewPost()

		// Concurrently viewing the post
		var wg sync.WaitGroup

		// Use multiple workers to view the post concurrently
		// this causes race condition, the count is not always correct as the workers can interrupt this
		// After mutex implementation each run should have the full count of the views for the post.
		workers := 5
		for i := 0; i < workers; i++ {
			wg.Go(
				func() {
					for range 1000 {
						post1.viewPost()
					}
				},
			)
		}

		wg.Wait()

		fmt.Printf("Run %d: %d (Expected 5001)\n", run+1, post1.postCount())
	}
}
