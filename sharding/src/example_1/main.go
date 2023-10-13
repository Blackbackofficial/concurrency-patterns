package main

import (
	"crypto/sha1"
	"fmt"
	"sync"
)

// Shard represents a shard with a read-write mutex and a map for storing data.
type Shard[V any] struct {
	sync.RWMutex
	m map[string]V
}

// ShardedMap is a slice of shards for sharding data.
type ShardedMap[V any] []*Shard[V]

// NewShardedMap creates and initializes a sharded map with the specified number of shards (nShards).
func NewShardedMap[V any](nShards int) ShardedMap[V] {
	shards := make([]*Shard[V], nShards)

	for i := range shards {
		shards[i] = &Shard[V]{m: make(map[string]V)}
	}

	return shards
}

// Get retrieves a value associated with the given key from the sharded map.
func (m ShardedMap[V]) Get(key string) V {
	shard := m.shard(key)
	shard.RLock()
	defer shard.RUnlock()

	return shard.m[key]
}

// Set associates a value with the given key in the sharded map.
func (m ShardedMap[V]) Set(key string, value V) {
	shard := m.shard(key)
	shard.Lock()
	defer shard.Unlock()

	shard.m[key] = value
}

// Keys retrieves all keys from the sharded map.
func (m ShardedMap[V]) Keys() []string {
	keys := make([]string, 0)

	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	wg.Add(len(m))

	for _, shard := range m {
		shard := shard
		go func() {
			defer wg.Done()
			shard.RLock()
			defer shard.RUnlock()

			for key := range shard.m {
				mu.Lock()
				keys = append(keys, key)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	return keys
}

// shardIndex calculates the shard index for a given key using SHA1 hashing.
func (m ShardedMap[V]) shardIndex(key string) int {
	chksm := sha1.Sum([]byte(key))
	hash := int(chksm[17])
	return hash % len(m)
}

// shard returns the shard associated with the given key.
func (m ShardedMap[V]) shard(key string) *Shard[V] {
	idx := m.shardIndex(key)
	return m[idx]
}

func main() {
	m := NewShardedMap[int](5)

	m.Set("alpha", 1)
	m.Set("beta", 2)
	m.Set("gamma", 3)

	fmt.Println(m.Get("alpha"))
	fmt.Println(m.Get("beta"))
	fmt.Println(m.Get("gamma"))

	keys := m.Keys()
	for _, k := range keys {
		fmt.Println(k)
	}
}
