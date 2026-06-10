package main

import (
	"fmt"
	"sync"
)

/*
* Mutex
* */

// Example of mutex using a struct post
type post struct {
	views int
}

// create an constructor function for the post struct
func createPost() post {
	p := post{views: 0}
	return p
}

// function which increments the post view count by 1
func (p *post) viewPost() {
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
