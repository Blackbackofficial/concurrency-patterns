package main

import (
	"fmt"
	"sync"
)

type Request struct {
	ID   int
	Data string
}

func main() {
	requests := []Request{
		{1, "Request A"},
		{2, "Request B"},
		{3, "Request C"},
		{4, "Request D"},
		{5, "Request E"},
		{6, "Request F"},
	}

	shards := make([][]Request, 3) // Creating 3 shards for balancing

	for i := range shards {
		shards[i] = make([]Request, 0)
	}

	var wg sync.WaitGroup

	for _, req := range requests {
		wg.Add(1)
		go func(req Request) {
			defer wg.Done()
			shardIndex := loadBalance(req.ID, len(shards)) // Selecting the shard index for balancing

			shards[shardIndex] = append(shards[shardIndex], req)
		}(req)
	}

	wg.Wait()

	// Output data from all shards
	for i, shard := range shards {
		fmt.Printf("Shard %d:\n", i)
		for _, req := range shard {
			fmt.Printf("Request ID: %d, Data: %s\n", req.ID, req.Data)
		}
	}
}

// Function for load balancing
func loadBalance(requestID, numShards int) int {
	return requestID % numShards
}
