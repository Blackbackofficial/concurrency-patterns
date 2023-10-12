package main

import (
	"fmt"
	"sync"
)

func mergeData(done <-chan struct{}, datas ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	merged := make(chan int)

	for _, data := range datas {
		data := data
		go func() {
			defer wg.Done()
			for {
				select {
				case val, ok := <-data:
					if !ok {
						return
					}
					select {
					case merged <- val:
					case <-done:
						return
					}
				case <-done:
					return
				}
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
	done := make(chan struct{})
	defer close(done)

	data1 := make(chan int)
	data2 := make(chan int)

	go func() {
		defer close(data1)
		for i := 1; i <= 5; i++ {
			data1 <- i
		}
	}()

	go func() {
		defer close(data2)
		for i := 6; i <= 10; i++ {
			data2 <- i
		}
	}()

	merged := mergeData(done, data1, data2)

	for val := range merged {
		fmt.Println(val)
	}
}
