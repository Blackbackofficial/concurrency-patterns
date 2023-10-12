package main

import (
	"fmt"
	"net/http"
)

func processRequest(url string, done chan<- RequestResult) {
	// Executing an HTTP request and processing the result
	resp, err := http.Get(url)
	if err != nil {
		done <- RequestResult{URL: url, Success: false, Error: err}
	} else {
		// Processing a successful request
		done <- RequestResult{URL: url, Success: true, Response: resp}
	}
}

func main() {
	urls := []string{"https://example.com", "https://example.org", "https://example.net"}

	done := make(chan RequestResult, len(urls))

	for _, url := range urls {
		go processRequest(url, done)
	}

	// Waiting for the completion of all queries and analyzing the results
	for i := 0; i < len(urls); i++ {
		result := <-done
		if result.Success {
			fmt.Printf("Request to %s succeeded\n", result.URL)
		} else {
			fmt.Printf("Request to %s failed with error: %v\n", result.URL, result.Error)
		}
	}
	close(done)
}

type RequestResult struct {
	URL      string
	Success  bool
	Response *http.Response
	Error    error
}
