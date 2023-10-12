package main

import "time"

func main() {
	result := make(chan interface{})
	done := make(chan struct{})

	go func() { // launch sub goroutine
		select {
		case result <- ...:  // write to result channel
		case <-done:
			// abort when `done` is closed
		}
	}()

	// wait for result or timeout, whichever occurs first
	select {
	case r := <-result:
		// process result
	case <-time.After(time.Second):
		// handle timeout
	}

	close(done) // cleanup `done` channel
}