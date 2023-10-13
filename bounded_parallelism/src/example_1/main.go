package main

import (
	"fmt"
	"sync"
	"time"
)

func Process(concurrency int, data []int) {
	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup
	wg.Add(len(data))

	for _, d := range data {
		sem <- struct{}{} // blocks until capacity is freed up

		go func(d int) {
			defer wg.Done()

			doProcess(d)
			<-sem // free up capacity
		}(d)
	}

	wg.Wait()
}

func main() {
	data := getData()

	Process(3, data)
}

func doProcess(d int) {
	time.Sleep(time.Second)
	fmt.Println("processed: ", d)
}

func getData() []int {
	data := make([]int, 10)

	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	return data
}
