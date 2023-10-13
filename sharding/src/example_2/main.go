package main

import (
	"fmt"
	"sync"
)

// Shard represents a shard (a partition) for storing data.
type Shard struct {
	sync.RWMutex
	data map[string]string
}

// ShardedStorage is a collection of shards for distributed data storage.
type ShardedStorage struct {
	shards []*Shard
}

// NewShardedStorage creates a new sharded storage with the specified number of shards.
func NewShardedStorage(numShards int) *ShardedStorage {
	shards := make([]*Shard, numShards)
	for i := 0; i < numShards; i++ {
		shards[i] = &Shard{data: make(map[string]string)}
	}
	return &ShardedStorage{shards: shards}
}

// Set adds a key-value pair to the sharded storage.
func (ss *ShardedStorage) Set(key, value string) {
	shard := ss.shards[shardIndex(key, len(ss.shards))]
	shard.Lock()
	defer shard.Unlock()
	shard.data[key] = value
}

// Get retrieves the value associated with a key from the sharded storage.
func (ss *ShardedStorage) Get(key string) (string, bool) {
	shard := ss.shards[shardIndex(key, len(ss.shards))]
	shard.RLock()
	defer shard.RUnlock()
	value, exists := shard.data[key]
	return value, exists
}

// shardIndex calculates the index of the shard for a given key.
func shardIndex(key string, numShards int) int {
	// A simple hashing function to determine the shard.
	// In a real-world scenario, a more sophisticated hashing function should be used.
	hash := hashString(key)
	return hash % numShards
}

// hashString is a simple hash function for demonstration purposes.
func hashString(s string) int {
	hash := 0
	for i := 0; i < len(s); i++ {
		hash += int(s[i])
	}
	return hash
}

func main() {
	numShards := 5
	shardedStorage := NewShardedStorage(numShards)

	keysAndValues := map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"carrot": "orange",
		"grape":  "purple",
		"lemon":  "yellow",
	}

	// Store data in the sharded storage.
	for key, value := range keysAndValues {
		shardedStorage.Set(key, value)
	}

	// Retrieve and print data from the sharded storage.
	for key, value := range keysAndValues {
		storedValue, exists := shardedStorage.Get(key)
		if exists {
			fmt.Printf("%s: %s (Stored: %s)\n", key, value, storedValue)
		} else {
			fmt.Printf("%s: %s (Not found)\n", key, value)
		}
	}
}
