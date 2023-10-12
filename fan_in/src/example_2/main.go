package main

import (
	"fmt"
	"sync"
)

func mergeChannels(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	merged := make(chan int)

	for _, ch := range channels {
		ch := ch
		go func() {
			defer wg.Done()
			for val := range ch {
				merged <- val
			}
		}()
		wg.Add(1)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 6; i <= 10; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(ch3)
		for i := 11; i <= 15; i++ {
			ch3 <- i
		}
	}()

	merged := mergeChannels(ch1, ch2, ch3)

	for val := range merged {
		fmt.Println(val)
	}
}
