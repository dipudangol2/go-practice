package main

import (
	"fmt"
	"sync"
	"time"
)

// Basic goroutine
func say(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(msg)
		time.Sleep(10 * time.Millisecond)
	}
}

// Channel communication
func producer(ch chan<- int, n int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)
}

// Fan-out: distribute work across multiple workers
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- j * j // square the number
	}
}

// fibonacci sends n fibonacci numbers on ch, then closes it.
func fibonacci(n int, ch chan<- int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	// --- WaitGroup: wait for goroutines to finish ---
	var wg sync.WaitGroup
	wg.Add(2)
	go say("ping", &wg)
	go say("pong", &wg)
	wg.Wait()
	fmt.Println("--- done with ping/pong ---")

	// --- Unbuffered channel ---
	ch := make(chan int)
	go producer(ch, 5)
	for v := range ch {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// --- Buffered channel ---
	bch := make(chan string, 3)
	bch <- "first"
	bch <- "second"
	bch <- "third"
	fmt.Println(<-bch, <-bch, <-bch)

	// --- Fan-out worker pool ---
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	var wg2 sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg2.Add(1)
		go worker(w, jobs, results, &wg2)
	}
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	wg2.Wait()
	close(results)

	squares := []int{}
	for r := range results {
		squares = append(squares, r)
	}
	fmt.Println("Squares:", squares)

	// --- Select + fibonacci ---
	fibCh := make(chan int, 10)
	fibs := []int{}
	go fibonacci(10, fibCh)
	for v := range fibCh {
		fibs = append(fibs, v)
	}
	fmt.Println("Fibonacci:", fibs)

	// --- Mutex for safe concurrent counter ---
	var mu sync.Mutex
	counter := 0
	var wg3 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg3.Wait()
	fmt.Println("Counter (should be 100):", counter)
}
