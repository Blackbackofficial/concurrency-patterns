package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sites := []string{"https://site1.com", "https://site2.com", "https://site3.com", "https://site4.com", "https://site5.com"}

	var wg sync.WaitGroup
	concurrencyLimit := 2
	sem := make(chan struct{}, concurrencyLimit)

	for _, site := range sites {
		sem <- struct{}{} // Ограничиваем параллельные запросы
		wg.Add(1)

		go func(site string) {
			defer func() {
				<-sem // Освобождаем слот
				wg.Done()
			}()

			data, err := fetchDataFromSite(site)
			if err != nil {
				fmt.Printf("Error fetching data from %s: %v\n", site, err)
				return
			}
			fmt.Printf("Data from %s: %s\n", site, data)
		}(site)
	}

	wg.Wait()
}

func fetchDataFromSite(url string) (string, error) {
	// Здесь выполняем HTTP-запрос и возвращаем данные
	time.Sleep(time.Second) // Имитируем запрос
	return "Data from " + url, nil
}
