package main

import (
	"fmt"
	"sync"
)

// FanOutWithWaitGroup функция выполняет Fan-Out, разделяя данные из источника src на n каналов.
// Возвращает слайс каналов, в которые данные будут отправлены.
func FanOutWithWaitGroup(src <-chan int, n int) []chan int {
	dests := make([]chan int, 0)

	// Создаем n каналов и горутин для каждого канала.
	for i := 0; i < n; i++ {
		dest := make(chan int)
		dests = append(dests, dest)

		// Горутина для чтения из src и отправки в канал dest.
		go func() {
			defer close(dest) // Закрываем канал при завершении горутины.

			for v := range src {
				dest <- v
			}
		}()
	}

	return dests
}

func main() {
	src := make(chan int)

	// Горутина для наполнения источника src данными.
	go func() {
		defer close(src)

		for i := 1; i <= 10; i++ {
			src <- i
		}
	}()

	n := 5 // Количество целевых каналов.

	dests := FanOutWithWaitGroup(src, n)

	var wg sync.WaitGroup
	wg.Add(len(dests))

	// Горутины для обработки данных из каждого целевого канала.
	for i, dest := range dests {
		go func(i int, dest <-chan int) {
			defer wg.Done()

			for v := range dest {
				fmt.Printf("Worker %d: %d\n", i, v)
			}
		}(i, dest)
	}

	// Ожидание завершения всех горутин.
	wg.Wait()
}
