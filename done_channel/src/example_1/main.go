package main

import (
	"fmt"
	"sync"
)

func worker(id int, done chan bool) {
	fmt.Printf("Worker %d is working\n", id)
	// Some work that goroutine does
	done <- true // Sending a completion notification
}

func main() {
	numWorkers := 5
	done := make(chan bool, numWorkers) // Creating a done channel

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Launching goroutines
	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			defer wg.Done() // Reducing the counter at completion
			worker(id, done)
		}(i)
	}

	// We are waiting for the completion of all goroutines
	go func() {
		wg.Wait()
		close(done) // Closing the channel when all goroutines have ended
	}()

	// Waiting until all goroutines are completed
	for range done {
		fmt.Println("Received a completion signal from a worker.")
	}
}
