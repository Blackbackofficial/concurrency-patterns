package main

import (
	"fmt"
	"sync"
)

func FanIn(sources ...<-chan int) <-chan int {
	dest := make(chan int) // return destination channel

	var wg sync.WaitGroup
	wg.Add(len(sources))

	// multiplex each source into dest
	for _, src := range sources {
		go func(src <-chan int) {
			defer wg.Done()

			for v := range src {
				dest <- v
			}
		}(src)
	}

	// close dest when all sources are closed
	go func() {
		wg.Wait()
		close(dest)
	}()

	return dest
}

func main() {
	src := sources()
	dest := FanIn(src...)

	for r := range dest {
		fmt.Println(r)
	}
}

func sources() []<-chan int {
	sources := make([]<-chan int, 0)

	for range []int{1, 2, 3} {
		src := make(chan int)
		sources = append(sources, src)

		go func() {
			defer close(src)

			for _, n := range []int{1, 2, 3, 4, 5} {
				src <- n
			}
		}()
	}

	return sources
}
