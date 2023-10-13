## Example 5:

This code is an extension of the previous example (Example 4) and continues to demonstrate the "Sharding" pattern for load balancing and request distribution. In this version, it places additional emphasis on thread safety to ensure that the code is robust against concurrent access. Here's how it works:

1. **Data Structure**: The `Request` struct represents individual requests with two fields, `ID` (an integer) and `Data` (a string).

2. **Main Function**:
   - An array of `Request` elements, `requests`, is created to represent the requests that need to be evenly balanced and distributed across shards.
   - A slice of slices of `Request`, `shards`, is created to simulate the shards. In this example, three shards are employed.
   - A `sync.WaitGroup` (`wg`) is used to ensure that all goroutines finish their work before proceeding.

3. **Goroutines for Request Distribution**:
   - For each `Request` in the `requests` array, a goroutine is initiated. Each goroutine is responsible for distributing a request to one of the three shards.
   - The `loadBalance` function calculates the shard index for a given request based on its `ID`. It employs a simple modulus operation to ensure that requests are evenly distributed among the three shards.

4. **Mutexes for Thread Safety**:
   - To fortify the code against potential data races and ensure thread safety, an array of `sync.Mutex` elements, `shardLocks`, is created. Each element corresponds to a specific shard and is used to synchronize access to that shard.
   - Before adding data to a shard, a goroutine acquires the corresponding mutex for that shard using `shardLocks[shardIndex].Lock()`.
   - After appending the request data to the shard, the goroutine promptly releases the mutex with `shardLocks[shardIndex].Unlock()`.

5. **Waiting for Goroutines**: The `main` function patiently waits for all goroutines to successfully complete their tasks, ensuring that every request is processed correctly, using the `wg.Wait()` statement.

6. **Outputting Requests from Shards**:
   - Once all requests have been appropriately distributed and processed, the `main` function proceeds to output the requests stored in each shard.
   - It systematically iterates through the `shards` slice and prints the request IDs and associated data within each shard.

This code prioritizes thread safety by incorporating mutexes, ensuring that concurrent access to the shards is handled securely. As a result, it is well-prepared to handle multiple goroutines simultaneously distributing requests to different shards without the risk of data races or other conflicts. This stands in contrast to "Example 4," where thread safety was not a consideration, potentially leading to concurrency issues when accessed concurrently.