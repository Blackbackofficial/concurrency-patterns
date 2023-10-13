package main

import (
	"fmt"
	"sync"
)

type Data struct {
	Key   string
	Value int
}

func main() {
	data := []Data{
		{"A", 1},
		{"B", 2},
		{"C", 3},
		{"D", 4},
		{"E", 5},
	}

	shards := make([]map[string]int, 3) // Creating 3 sharja

	for i := range shards {
		shards[i] = make(map[string]int)
	}

	var wg sync.WaitGroup

	for _, d := range data {
		wg.Add(1)
		go func(d Data) {
			defer wg.Done()
			shard := shards[hashCode(d.Key)%3] // Distributing data between shards

			shard[d.Key] = d.Value
		}(d)
	}

	wg.Wait()

	// Output data from all shards
	for i, shard := range shards {
		fmt.Printf("Shard %d:\n", i)
		for k, v := range shard {
			fmt.Printf("Key: %s, Value: %d\n", k, v)
		}
	}
}

// Function for key hashing and shard distribution
func hashCode(key string) int {
	// In real applications, more complex hashing should be used
	// For the simplicity of the example, we use a simple hash
	hash := 0
	for _, c := range key {
		hash += int(c)
	}
	return hash
}
