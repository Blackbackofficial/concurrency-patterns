package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func FanOutWithCancel(ctx context.Context, src <-chan int, n int) []<-chan int {
	dests := make([]<-chan int, 0)

	for i := 0; i < n; i++ {
		dest := make(chan int)
		dests = append(dests, dest)

		go func() {
			defer close(dest)

			for {
				select {
				case v, ok := <-src:
					if !ok {
						return
					}
					select {
					case dest <- v:
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	return dests
}

func source() <-chan int {
	src := make(chan int)

	go func() {
		defer close(src)

		for i := 1; i <= 10; i++ {
			src <- i
		}
	}()

	return src
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	src := source()
	dests := FanOutWithCancel(ctx, src, 5)

	var wg sync.WaitGroup
	wg.Add(len(dests))

	for i, dest := range dests {
		go func(i int, dest <-chan int) {
			defer wg.Done()

			for v := range dest {
				fmt.Printf("Worker %d: %d\n", i, v)
			}
		}(i, dest)
	}

	// Simulate some work
	time.Sleep(1 * time.Second)

	// Cancel the context to stop the Fan-Out
	cancel()
	wg.Wait()
}
