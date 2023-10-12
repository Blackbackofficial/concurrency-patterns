package main

import (
	"fmt"
	"sync"
)

func FanOut(src <-chan int, n int) []<-chan int {
	dests := make([]<-chan int, 0)

	for i := 0; i < n; i++ {
		dest := make(chan int)
		dests = append(dests, dest)

		go func() {
			defer close(dest)

			for v := range src {
				dest <- v
			}
		}()
	}

	return dests
}

func main() {
	src := source()
	dests := FanOut(src, 5)

	var wg sync.WaitGroup
	wg.Add(len(dests))

	for _, dest := range dests {
		go func(dest <-chan int) {
			defer wg.Done()

			for v := range dest {
				fmt.Println(v)
			}
		}(dest)
	}

	wg.Wait()
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
