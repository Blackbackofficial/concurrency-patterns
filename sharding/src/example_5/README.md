## Example 5:

This code is similar to the previous example (example 4), demonstrating the "Sharding" pattern for load balancing and request distribution. In this version, there is an additional focus on thread safety to ensure that the code is safe for concurrent access. Here's how it works:

1. **Data Structure**: The `Request` struct represents individual requests with an `ID` (an integer) and `Data` (a string).

2. **Main Function**:
    - An array of `Request` elements, `requests`, is created to represent the requests that need to be balanced and distributed across shards.
    - A slice of slices of `Request`, `shards`, is created to simulate the shards. In this case, three shards are created.
    - A `sync.WaitGroup` (`wg`) is used to ensure that all goroutines finish their work before proceeding.

3. **Goroutines for Request Distribution**:
    - For each `Request` in the `requests` array, a goroutine is started. Each goroutine is responsible for distributing a request to one of the three shards.
    - The `loadBalance` function calculates the shard index for a given request based on its `ID`, using a simple modulus operation to evenly distribute requests among the three shards.

4. **Mutexes for Thread Safety**:
    - To make the code thread-safe, an array of `sync.Mutex` elements, `shardLocks`, is created. Each element corresponds to a shard and is used to synchronize access to that shard.
    - Before adding data to a shard, a goroutine captures the mutex for that shard using `shardLocks[shardIndex].Lock()`.
    - After appending the request data to the shard, the goroutine releases the mutex with `shardLocks[shardIndex].Unlock()`.

5. **Waiting for Goroutines**: The `main` function waits for all goroutines to finish their work using the `wg.Wait()` statement.

6. **Outputting Requests from Shards**:
    - After all requests are distributed, the `main` function proceeds to output the requests from all shards.
    - It iterates through the `shards` slice and prints the request IDs and data within each shard.

This code emphasizes thread safety by using mutexes to protect access to individual shards. As a result, it's designed to be safe for concurrent access, ensuring that multiple goroutines can safely distribute requests to different shards without conflicts. This is in contrast to "Example 4," where thread safety was not considered, potentially leading to data races and issues when accessed concurrently.